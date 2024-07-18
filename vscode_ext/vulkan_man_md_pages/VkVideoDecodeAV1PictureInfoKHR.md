# VkVideoDecodeAV1PictureInfoKHR(3) Manual Page

## Name

VkVideoDecodeAV1PictureInfoKHR - Structure specifies AV1 picture
information when decoding a frame



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkVideoDecodeAV1PictureInfoKHR` structure is defined as:

``` c
// Provided by VK_KHR_video_decode_av1
typedef struct VkVideoDecodeAV1PictureInfoKHR {
    VkStructureType                        sType;
    const void*                            pNext;
    const StdVideoDecodeAV1PictureInfo*    pStdPictureInfo;
    int32_t                                referenceNameSlotIndices[VK_MAX_VIDEO_AV1_REFERENCES_PER_FRAME_KHR];
    uint32_t                               frameHeaderOffset;
    uint32_t                               tileCount;
    const uint32_t*                        pTileOffsets;
    const uint32_t*                        pTileSizes;
} VkVideoDecodeAV1PictureInfoKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `pStdPictureInfo` is a pointer to a `StdVideoDecodeAV1PictureInfo`
  structure specifying <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#decode-av1-picture-info"
  target="_blank" rel="noopener">AV1 picture information</a>.

- `referenceNameSlotIndices` is an array of seven
  (`VK_MAX_VIDEO_AV1_REFERENCES_PER_FRAME_KHR`, which is equal to the
  Video Std definition `STD_VIDEO_AV1_REFS_PER_FRAME`) signed integer
  values specifying the index of the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#dpb-slot"
  target="_blank" rel="noopener">DPB slot</a> or a negative integer
  value for each <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#decode-av1-reference-names"
  target="_blank" rel="noopener">AV1 reference name</a> used for inter
  coding. In particular, the DPB slot index for the AV1 reference name
  `frame` is specified in `referenceNameSlotIndices`\[`frame` -
  `STD_VIDEO_AV1_REFERENCE_NAME_LAST_FRAME`\].

- `frameHeaderOffset` is the byte offset of the AV1 frame header OBU, as
  defined in section 5.9 of the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#aomedia-av1"
  target="_blank" rel="noopener">AV1 Specification</a>, within the video
  bitstream buffer range specified in
  [VkVideoDecodeInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeInfoKHR.html).

- `tileCount` is the number of elements in `pTileOffsets` and
  `pTileSizes`.

- `pTileOffsets` is a pointer to an array of `tileCount` integers
  specifying the byte offset of the tiles of the picture within the
  video bitstream buffer range specified in
  [VkVideoDecodeInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeInfoKHR.html).

- `pTileSizes` is a pointer to an array of `tileCount` integers
  specifying the byte size of the tiles of the picture within the video
  bitstream buffer range specified in
  [VkVideoDecodeInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeInfoKHR.html).

## <a href="#_description" class="anchor"></a>Description

This structure is specified in the `pNext` chain of the
[VkVideoDecodeInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeInfoKHR.html) structure passed to
[vkCmdDecodeVideoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdDecodeVideoKHR.html) to specify the
codec-specific picture information for an <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#decode-av1"
target="_blank" rel="noopener">AV1 decode operation</a>.

Decode Output Picture Information  
When this structure is specified in the `pNext` chain of the
[VkVideoDecodeInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeInfoKHR.html) structure passed to
[vkCmdDecodeVideoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdDecodeVideoKHR.html), the information related
to the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#decode-output-picture-info"
target="_blank" rel="noopener">decode output picture</a> is defined as
follows:

- The image subregion used is determined according to the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#decode-av1-picture-data-access"
  target="_blank" rel="noopener">AV1 Decode Picture Data Access</a>
  section.

- The decode output picture is associated with the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#decode-av1-picture-info"
  target="_blank" rel="noopener">AV1 picture information</a> provided in
  `pStdPictureInfo`.

<!-- -->

Std Picture Information  
The members of the `StdVideoDecodeAV1PictureInfo` structure pointed to
by `pStdPictureInfo` are interpreted as follows:

- `flags.reserved`, `reserved1`, and `reserved2` are used only for
  padding purposes and are otherwise ignored;

- <span id="decode-av1-film-grain"></span> `flags.apply_grain` indicates
  that film grain is enabled for the decoded picture, as defined in
  section 6.8.20 of the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#aomedia-av1"
  target="_blank" rel="noopener">AV1 Specification</a>;

- `tg_start` and `tg_end` are interpreted as defined in section 6.10.1
  of the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#aomedia-av1"
  target="_blank" rel="noopener">AV1 Specification</a>;

- `OrderHint`, `OrderHints`, and `expectedFrameId` are interpreted as
  defined in section 6.8.2 of the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#aomedia-av1"
  target="_blank" rel="noopener">AV1 Specification</a>;

- the `StdVideoAV1TileInfo` structure pointed to by `pTileInfo` is
  interpreted as follows:

  - `flags.reserved` and `reserved1` are used only for padding purposes
    and are otherwise ignored;

  - `pMiColStarts` is a pointer to an array of `TileCols` number of
    unsigned integers that corresponds to `MiColStarts` defined in
    section 6.8.14 of the <a
    href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#aomedia-av1"
    target="_blank" rel="noopener">AV1 Specification</a>;

  - `pMiRowStarts` is a pointer to an array of `TileRows` number of
    unsigned integers that corresponds to `MiRowStarts` defined in
    section 6.8.14 of the <a
    href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#aomedia-av1"
    target="_blank" rel="noopener">AV1 Specification</a>;

  - `pWidthInSbsMinus1` is a pointer to an array of `TileCols` number of
    unsigned integers that corresponds to `width_in_sbs_minus_1` defined
    in section 6.8.14 of the <a
    href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#aomedia-av1"
    target="_blank" rel="noopener">AV1 Specification</a>;

  - `pHeightInSbsMinus1` is a pointer to an array of `TileRows` number
    of unsigned integers that corresponds to `height_in_sbs_minus_1`
    defined in section 6.8.14 of the <a
    href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#aomedia-av1"
    target="_blank" rel="noopener">AV1 Specification</a>;

  - all other members of `StdVideoAV1TileInfo` are interpreted as
    defined in section 6.8.14 of the <a
    href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#aomedia-av1"
    target="_blank" rel="noopener">AV1 Specification</a>;

- the `StdVideoAV1Quantization` structure pointed to by `pQuantization`
  is interpreted as follows:

  - `flags.reserved` is used only for padding purposes and is otherwise
    ignored;

  - all other members of `StdVideoAV1Quantization` are interpreted as
    defined in section 6.8.11 of the <a
    href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#aomedia-av1"
    target="_blank" rel="noopener">AV1 Specification</a>;

- if `flags.segmentation_enabled` is set, then the
  `StdVideoAV1Segmentation` structure pointed to by `pSegmentation` is
  interpreted as follows:

  - the elements of `FeatureEnabled` are bitmasks where bit index j of
    element i corresponds to `FeatureEnabled[i][j]` as defined in
    section 6.8.13 of the <a
    href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#aomedia-av1"
    target="_blank" rel="noopener">AV1 Specification</a>;

  - `FeatureData` is interpreted as defined in section 6.8.13 of the <a
    href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#aomedia-av1"
    target="_blank" rel="noopener">AV1 Specification</a>;

- the `StdVideoAV1LoopFilter` structure pointed to by `pLoopFilter` is
  interpreted as follows:

  - `flags.reserved` is used only for padding purposes and is otherwise
    ignored;

  - `update_ref_delta` is a bitmask where bit index i is interpreted as
    the value of `update_ref_delta` corresponding to element i of
    `loop_filter_ref_deltas` as defined in section 6.8.10 of the <a
    href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#aomedia-av1"
    target="_blank" rel="noopener">AV1 Specification</a>;

  - `update_mode_delta` is a bitmask where bit index i is interpreted as
    the value of `update_mode_delta` corresponding to element i of
    `loop_filter_mode_deltas` as defined in section 6.8.10 of the <a
    href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#aomedia-av1"
    target="_blank" rel="noopener">AV1 Specification</a>;

  - all other members of `StdVideoAV1LoopFilter` are interpreted as
    defined in section 6.8.10 of the <a
    href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#aomedia-av1"
    target="_blank" rel="noopener">AV1 Specification</a>;

- if `flags.enable_cdef` is set in the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#decode-av1-active-sequence-header"
  target="_blank" rel="noopener">active sequence header</a>, then the
  members of the `StdVideoAV1CDEF` structure pointed to by `pCDEF` are
  interpreted as follows:

  - `cdef_y_sec_strength` and `cdef_uv_sec_strength` are the bitstream
    values of the corresponding syntax elements defined in section
    5.9.19 of the <a
    href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#aomedia-av1"
    target="_blank" rel="noopener">AV1 Specification</a>;

  - all other members of `StdVideoAV1CDEF` are interpreted as defined in
    section 6.10.14 of the <a
    href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#aomedia-av1"
    target="_blank" rel="noopener">AV1 Specification</a>;

- the `StdVideoAV1LoopRestoration` structure pointed to by
  `pLoopRestoration` is interpreted as follows:

  - `LoopRestorationSize`\[`plane`\] is interpreted as log2(`size`) - 5,
    where `size` is the value of `LoopRestorationSize`\[`plane`\] as
    defined in section 6.10.15 of the <a
    href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#aomedia-av1"
    target="_blank" rel="noopener">AV1 Specification</a>.

  - all other members of `StdVideoAV1LoopRestoration` are defined as in
    section 6.10.15 of the <a
    href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#aomedia-av1"
    target="_blank" rel="noopener">AV1 Specification</a>;

- the members of the `StdVideoAV1GlobalMotion` structure provided in
  `global_motion` are interpreted as defined in section 7.10 of the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#aomedia-av1"
  target="_blank" rel="noopener">AV1 Specification</a>;

- if `flags.film_grain_params_present` is set in the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#decode-av1-active-sequence-header"
  target="_blank" rel="noopener">active sequence header</a>, then the
  `StdVideoAV1FilmGrain` structure pointed to by `pFilmGrain` is
  interpreted as follows:

  - `flags.reserved` is used only for padding purposes and is otherwise
    ignored;

  - all other members of `StdVideoAV1FilmGrain` are interpreted as
    defined in section 6.8.20 of the <a
    href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#aomedia-av1"
    target="_blank" rel="noopener">AV1 Specification</a>;

- all other members are interpreted as defined in section 6.8 of the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#aomedia-av1"
  target="_blank" rel="noopener">AV1 Specification</a>.

When <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#decode-av1-film-grain"
target="_blank" rel="noopener">film grain is enabled</a> for the decoded
frame, the `flags.update_grain` and `film_grain_params_ref_idx` values
specified in `StdVideoAV1FilmGrain` are ignored by AV1 decode operations
and the `load_grain_params` function, as defined in section 6.8.20 of
the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#aomedia-av1"
target="_blank" rel="noopener">AV1 Specification</a>, is not executed.
Instead, the application is responsible for specifying the effective
film grain parameters for the frame in `StdVideoAV1FilmGrain`.

When <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#decode-av1-film-grain"
target="_blank" rel="noopener">film grain is enabled</a> for the decoded
frame, the application is required to specify a different decode output
picture resource in
[VkVideoDecodeInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeInfoKHR.html)::`dstPictureResource`
compared to the reconstructed picture specified in
[VkVideoDecodeInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeInfoKHR.html)::`pSetupReferenceSlot->pPictureResource`
even if the implementation does not report support for
`VK_VIDEO_DECODE_CAPABILITY_DPB_AND_OUTPUT_DISTINCT_BIT_KHR` in
[VkVideoDecodeCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeCapabilitiesKHR.html)::`flags`
for the video decode profile.

Reference picture setup is controlled by the value of
`StdVideoDecodeAV1PictureInfo`::`refresh_frame_flags`. If it is not zero
and a <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#decode-reconstructed-picture-info"
target="_blank" rel="noopener">reconstructed picture</a> is specified,
then the latter is used as the target of picture reconstruction to <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#dpb-slot-states"
target="_blank" rel="noopener">activate</a> the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#dpb-slot"
target="_blank" rel="noopener">DPB slot</a> specified in
`pDecodeInfo->pSetupReferenceSlot->slotIndex`. If
`StdVideoDecodeAV1PictureInfo`::`refresh_frame_flags` is zero, but a <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#decode-reconstructed-picture-info"
target="_blank" rel="noopener">reconstructed picture</a> is specified,
then the corresponding picture reference associated with the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#dpb-slot"
target="_blank" rel="noopener">DPB slot</a> is invalidated, as described
in the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#dpb-slot-states"
target="_blank" rel="noopener">DPB Slot States</a> section.

Active Parameter Sets  
The *active sequence header* is the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#decode-av1-sequence-header"
target="_blank" rel="noopener">AV1 sequence header</a> stored in the
bound video session parameters object.

Valid Usage (Implicit)

- <a href="#VUID-VkVideoDecodeAV1PictureInfoKHR-sType-sType"
  id="VUID-VkVideoDecodeAV1PictureInfoKHR-sType-sType"></a>
  VUID-VkVideoDecodeAV1PictureInfoKHR-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_VIDEO_DECODE_AV1_PICTURE_INFO_KHR`

- <a href="#VUID-VkVideoDecodeAV1PictureInfoKHR-pStdPictureInfo-parameter"
  id="VUID-VkVideoDecodeAV1PictureInfoKHR-pStdPictureInfo-parameter"></a>
  VUID-VkVideoDecodeAV1PictureInfoKHR-pStdPictureInfo-parameter  
  `pStdPictureInfo` **must** be a valid pointer to a valid
  `StdVideoDecodeAV1PictureInfo` value

- <a href="#VUID-VkVideoDecodeAV1PictureInfoKHR-pTileOffsets-parameter"
  id="VUID-VkVideoDecodeAV1PictureInfoKHR-pTileOffsets-parameter"></a>
  VUID-VkVideoDecodeAV1PictureInfoKHR-pTileOffsets-parameter  
  `pTileOffsets` **must** be a valid pointer to an array of `tileCount`
  `uint32_t` values

- <a href="#VUID-VkVideoDecodeAV1PictureInfoKHR-pTileSizes-parameter"
  id="VUID-VkVideoDecodeAV1PictureInfoKHR-pTileSizes-parameter"></a>
  VUID-VkVideoDecodeAV1PictureInfoKHR-pTileSizes-parameter  
  `pTileSizes` **must** be a valid pointer to an array of `tileCount`
  `uint32_t` values

- <a href="#VUID-VkVideoDecodeAV1PictureInfoKHR-tileCount-arraylength"
  id="VUID-VkVideoDecodeAV1PictureInfoKHR-tileCount-arraylength"></a>
  VUID-VkVideoDecodeAV1PictureInfoKHR-tileCount-arraylength  
  `tileCount` **must** be greater than `0`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_video_decode_av1](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_video_decode_av1.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkVideoDecodeAV1PictureInfoKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
