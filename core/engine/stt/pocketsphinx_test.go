package stt

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

const (
	hmm_path, hmmerr  = filepath.Abs("data/stt/models/en-us/en-us")
	dict_path, dicterr = filepath.Abs("data/stt/models/en-us/cmudict-en-us.dict")
	lm_path, lmerr   = filepath.Abs("data/stt/models/en-us/en-us.lm.bin")
)

func TestProcessUTT(t *testing.T) {
	conf := Config{
		Hmm:         hmm_path,
		Dict:        dict_path,
		Lm:          lm_path,
		DisableInfo: true,
	}

	ps := NewPocketSphinx(conf)
	defer ps.Free()
	ps.StartStream()
	file, fileerr := filepath.Abs("data/stt/speech/goforward.raw")
	if fileerr != nil {
		t.Error(fileerr)
		return
	}
	dat, err := ioutil.ReadFile(file)
	results, err := ps.ProcessUtt(dat, 2, true)
	if err != nil {
		t.Error(err)
		return
	}
	if len(results) == 0 {
		t.Error(results)
		return
	}
	if results[0].Text != "go forward ten meters" {
		t.Error("could not recognize", results[0].Text)
	}
}

func TestProcessRaw(t *testing.T) {
	conf := Config{
		Hmm:  hmm_path,
		Dict: dict_path,
		Lm:   lm_path,
		//Beam:        FloatParam(1e-400),
		//Wbeam:       FloatParam(1e-400),
		//Bestpath:    "no",
		DisableInfo: true,
	}

	ps := NewPocketSphinx(conf)
	defer ps.Free()
	ps.StartStream()
	file, fileerr := filepath.Abs("data/stt/speech/goforward.raw")
	if fileerr != nil {
		t.Error(fileerr)
		return
	}
	dat, _ := ioutil.ReadFile(file)
	err := ps.StartUtt()
	if err != nil {
		t.Error(err)
	}
	err = ps.ProcessRaw(dat, false, false)
	if err != nil {
		t.Error(err)
	}
	err = ps.EndUtt()
	if err != nil {
		t.Error(err)
	}
	r, err := ps.GetHyp(false)
	if err != nil {
		t.Error(err)
	}
	if r.Text != "go somewhere and do something" {
		t.Error("could not recognize", r.Text, r.Score)
	}
	dat, _ = ioutil.ReadFile(file)
	err = ps.StartUtt()
	if err != nil {
		t.Error(err)
	}
	err = ps.ProcessRaw(dat, false, true)
	if err != nil {
		t.Error(err)
	}
	err = ps.EndUtt()
	if err != nil {
		t.Error(err)
	}
	r, _ = ps.GetHyp(true)

	if r.Text != "go forward ten meters" {
		t.Error("could not recognize", r.Text, r.Score)
	}

}

func TestProcessRawIncremental(t *testing.T) {

	conf := Config{
		Hmm:         hmm_path,
		Dict:        dict_path,
		Lm:          lm_path,
		DisableInfo: true,
	}
	ps := NewPocketSphinx(conf)
	defer ps.Free()
	filepath, file_err := filepath.Abs("data/stt/speech/goforward.raw")
	if file_err != nil {
		t.Error(file_err)
		return
	}
	file, _ := os.Open(filepath)
	defer file.Close()
	buf := make([]byte, 1024)
	var lastr Result
	ps.StartUtt()
	for {
		size, err := file.Read(buf)
		if err != nil {
			break
		}
		ps.ProcessRaw(buf[:size], false, false)
		r, err := ps.GetHyp(false)
		if err == nil {
			lastr = r
		}
	}
	ps.EndUtt()
	if lastr.Text != "go forward ten meters" {
		t.Error("recognition failed")
	}
}

func TestWordSpottingUtt(t *testing.T) {
	conf := Config{
		Hmm:         hmm_path,
		Dict:        dict_path,
		Keyphrase:   "forward",
		DisableInfo: true,
	}

	ps := NewPocketSphinx(conf)
	defer ps.Free()
	file, fileerr := filepath.Abs("data/stt/speech/goforward.raw")
	if fileerr != nil {
		t.Error(fileerr)
		return
	}
	dat, _ := ioutil.ReadFile(file)
	ps.StartUtt()
	ps.ProcessRaw(dat, false, true)
	ps.EndUtt()
	r, _ := ps.GetHyp(false)
	if r.Text != "forward" {
		t.Error("could not recognize", r.Text)
	}
}
func TestWordSpotting(t *testing.T) {
	conf := Config{
		Hmm:         hmm_path,
		Dict:        dict_path,
		Keyphrase:   "forward",
		DisableInfo: true,
	}
	ps := NewPocketSphinx(conf)
	defer ps.Free()
	filepath, fileerr := filepath.Abs("data/stt/speech/goforward.raw")
	if fileerr != nil {
		t.Error(fileerr)
		return
	}
	file, _ := os.Open(filepath)
	defer file.Close()
	c := 0
	buf := make([]byte, 1024)
	ps.StartUtt()
	for {
		size, err := file.Read(buf)
		if err != nil {
			break
		}
		ps.ProcessRaw(buf[:size], false, false)
		r, err := ps.GetHyp(false)
		if err == nil && r.Text == "forward" {
			c += 1
			ps.EndUtt()
			ps.StartUtt()
		}
	}
	ps.EndUtt()
	if c == 0 {
		t.Error("keyphrase not found")
	}
}

func TestEndUttErr(t *testing.T) {
	conf := Config{
		Hmm:         hmm_path,
		Dict:        dict_path,
		Keyphrase:   "forward",
		DisableInfo: true,
	}
	ps := NewPocketSphinx(conf)
	defer ps.Free()
	if ps.EndUtt() == nil {
		t.Error("call EndUtt befoer StartUtt must raise error")
	}
}

func TestStartUttErr(t *testing.T) {
	conf := Config{
		Hmm:         hmm_path,
		Dict:        dict_path,
		Keyphrase:   "forward",
		DisableInfo: true,
	}
	ps := NewPocketSphinx(conf)
	defer ps.Free()
	if err := ps.StartUtt(); err != nil {
		t.Error(err)
		return
	}
	if ps.StartUtt() == nil {
		t.Error("call StartUtt after calling StartUtt without ErrUtt, must raise error")
	}
}
