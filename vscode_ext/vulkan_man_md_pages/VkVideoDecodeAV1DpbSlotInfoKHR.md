# VkVideoDecodeAV1DpbSlotInfoKHR(3) Manual Page

## Name

VkVideoDecodeAV1DpbSlotInfoKHR - Structure specifies AV1 DPB information
when decoding a frame



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkVideoDecodeAV1DpbSlotInfoKHR` structure is defined as:

``` c
// Provided by VK_KHR_video_decode_av1
typedef struct VkVideoDecodeAV1DpbSlotInfoKHR {
    VkStructureType                          sType;
    const void*                              pNext;
    const StdVideoDecodeAV1ReferenceInfo*    pStdReferenceInfo;
} VkVideoDecodeAV1DpbSlotInfoKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `pStdReferenceInfo` is a pointer to a `StdVideoDecodeAV1ReferenceInfo`
  structure specifying <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#decode-av1-reference-info"
  target="_blank" rel="noopener">AV1 reference information</a>.

## <a href="#_description" class="anchor"></a>Description

This structure is specified in the `pNext` chain of
[VkVideoDecodeInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeInfoKHR.html)::`pSetupReferenceSlot`,
if not `NULL`, and the `pNext` chain of the elements of
[VkVideoDecodeInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeInfoKHR.html)::`pReferenceSlots` to
specify the codec-specific reference picture information for an <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#decode-av1"
target="_blank" rel="noopener">AV1 decode operation</a>.

Active Reference Picture Information  
When this structure is specified in the `pNext` chain of the elements of
[VkVideoDecodeInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeInfoKHR.html)::`pReferenceSlots`,
one element is added to the list of <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#decode-active-reference-picture-info"
target="_blank" rel="noopener">active reference pictures</a> used by the
video decode operation for each element of
[VkVideoDecodeInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeInfoKHR.html)::`pReferenceSlots` as
follows:

- The image subregion used is determined according to the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#decode-av1-picture-data-access"
  target="_blank" rel="noopener">AV1 Decode Picture Data Access</a>
  section.

- The reference picture is associated with the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#dpb-slot"
  target="_blank" rel="noopener">DPB slot</a> index specified in the
  `slotIndex` member of the corresponding element of
  [VkVideoDecodeInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeInfoKHR.html)::`pReferenceSlots`.

- The reference picture is associated with the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#decode-av1-reference-info"
  target="_blank" rel="noopener">AV1 reference information</a> provided
  in `pStdReferenceInfo`.

<!-- -->

Reconstructed Picture Information  
When this structure is specified in the `pNext` chain of
[VkVideoDecodeInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeInfoKHR.html)::`pSetupReferenceSlot`,
the information related to the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#decode-reconstructed-picture-info"
target="_blank" rel="noopener">reconstructed picture</a> is defined as
follows:

- The image subregion used is determined according to the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#decode-av1-picture-data-access"
  target="_blank" rel="noopener">AV1 Decode Picture Data Access</a>
  section.

- If <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#decode-av1-ref-pic-setup"
  target="_blank" rel="noopener">reference picture setup</a> is
  requested, then the reconstructed picture is used to <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#dpb-slot-states"
  target="_blank" rel="noopener">activate</a> the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#dpb-slot"
  target="_blank" rel="noopener">DPB slot</a> with the index specified
  in
  [VkVideoDecodeInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeInfoKHR.html)::`pSetupReferenceSlot->slotIndex`.

- The reconstructed picture is associated with the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#decode-av1-reference-info"
  target="_blank" rel="noopener">AV1 reference information</a> provided
  in `pStdReferenceInfo`.

<!-- -->

Std Reference Information  
The members of the `StdVideoDecodeAV1ReferenceInfo` structure pointed to
by `pStdReferenceInfo` are interpreted as follows:

- `flags.reserved` and `reserved1` are used only for padding purposes
  and are otherwise ignored;

- `flags.disable_frame_end_update_cdf` is interpreted as defined in
  section 6.8.2 of the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#aomedia-av1"
  target="_blank" rel="noopener">AV1 Specification</a>;

- `flags.segmentation_enabled` is interpreted as defined in section
  6.8.13 of the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#aomedia-av1"
  target="_blank" rel="noopener">AV1 Specification</a>;

- `frame_type` is interpreted as defined in section 6.8.2 of the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#aomedia-av1"
  target="_blank" rel="noopener">AV1 Specification</a>;

  <table>
  <colgroup>
  <col style="width: 50%" />
  <col style="width: 50%" />
  </colgroup>
  <tbody>
  <tr>
  <td class="icon"><em></em></td>
  <td class="content">Note
  <p>The <code>frame_type</code> member is defined with the type
  <code>uint8_t</code>, but it takes the same values defined in the
  <code>StdVideoAV1FrameType</code> enumeration type as
  <code>StdVideoDecodeAV1PictureInfo</code>::<code>frame_type</code>.</p></td>
  </tr>
  </tbody>
  </table>

- `RefFrameSignBias` is a bitmask where bit index i corresponds to
  `RefFrameSignBias[i]` as defined in section 6.8.2 of the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#aomedia-av1"
  target="_blank" rel="noopener">AV1 Specification</a>;

- `OrderHint` is interpreted as defined in section 6.8.2 of the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#aomedia-av1"
  target="_blank" rel="noopener">AV1 Specification</a>;

- `SavedOrderHints` is interpreted as defined in section 7.20 of the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#aomedia-av1"
  target="_blank" rel="noopener">AV1 Specification</a>.

  <table>
  <colgroup>
  <col style="width: 50%" />
  <col style="width: 50%" />
  </colgroup>
  <tbody>
  <tr>
  <td class="icon"><em></em></td>
  <td class="content">Note
  <p>When the AV1 reference information is provided for the reconstructed
  picture, certain parameters (e.g. <code>frame_type</code>) are specified
  both in the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#decode-av1-picture-info"
  target="_blank" rel="noopener">AV1 picture information</a> and in the
  AV1 reference information. This is necessary because unlike the AV1
  picture information, which is only used for the purposes of the video
  decode operation in question, the AV1 reference information specified
  for the reconstructed picture <strong>may</strong> be associated with
  the activated DPB slot, meaning that some implementations
  <strong>may</strong> maintain it as part of the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#reference-metadata"
  target="_blank" rel="noopener">reference picture metadata</a>
  corresponding to the video picture resource associated with the DPB
  slot. When the AV1 reference information is provided for an active
  reference picture, the specified parameters correspond to the parameters
  specified when the DPB slot was activated (set up) with the reference
  picture, as usual, in order to communicate these parameters for
  implementations that do not maintain any subset of these parameters as
  part of the DPB slot’s <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#reference-metadata"
  target="_blank" rel="noopener">reference picture metadata</a>.</p></td>
  </tr>
  </tbody>
  </table>

Valid Usage (Implicit)

- <a href="#VUID-VkVideoDecodeAV1DpbSlotInfoKHR-sType-sType"
  id="VUID-VkVideoDecodeAV1DpbSlotInfoKHR-sType-sType"></a>
  VUID-VkVideoDecodeAV1DpbSlotInfoKHR-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_VIDEO_DECODE_AV1_DPB_SLOT_INFO_KHR`

- <a
  href="#VUID-VkVideoDecodeAV1DpbSlotInfoKHR-pStdReferenceInfo-parameter"
  id="VUID-VkVideoDecodeAV1DpbSlotInfoKHR-pStdReferenceInfo-parameter"></a>
  VUID-VkVideoDecodeAV1DpbSlotInfoKHR-pStdReferenceInfo-parameter  
  `pStdReferenceInfo` **must** be a valid pointer to a valid
  `StdVideoDecodeAV1ReferenceInfo` value

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_video_decode_av1](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_video_decode_av1.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkVideoDecodeAV1DpbSlotInfoKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
