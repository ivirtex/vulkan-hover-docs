# VkVideoDecodeAV1InlineSessionParametersInfoKHR(3) Manual Page

## Name

VkVideoDecodeAV1InlineSessionParametersInfoKHR - Structure specifies inline AV1 decoder parameter set information



## [](#_c_specification)C Specification

The `VkVideoDecodeAV1InlineSessionParametersInfoKHR` structure is defined as:

```c++
// Provided by VK_KHR_video_decode_av1 with VK_KHR_video_maintenance2
typedef struct VkVideoDecodeAV1InlineSessionParametersInfoKHR {
    VkStructureType                     sType;
    const void*                         pNext;
    const StdVideoAV1SequenceHeader*    pStdSequenceHeader;
} VkVideoDecodeAV1InlineSessionParametersInfoKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `pStdSequenceHeader` is `NULL` or a pointer to an instance of the `StdVideoAV1SequenceHeader` structure describing the [active AV1 sequence header](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#decode-av1-active-sequence-header).

## [](#_description)Description

If `pStdSequenceHeader` is not `NULL`, the issued video decode operations will use the specified sequence header parameters instead of the active sequence header being sourced from the bound video session parameters object.

Valid Usage (Implicit)

- [](#VUID-VkVideoDecodeAV1InlineSessionParametersInfoKHR-sType-sType)VUID-VkVideoDecodeAV1InlineSessionParametersInfoKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_VIDEO_DECODE_AV1_INLINE_SESSION_PARAMETERS_INFO_KHR`
- [](#VUID-VkVideoDecodeAV1InlineSessionParametersInfoKHR-pStdSequenceHeader-parameter)VUID-VkVideoDecodeAV1InlineSessionParametersInfoKHR-pStdSequenceHeader-parameter  
  If `pStdSequenceHeader` is not `NULL`, `pStdSequenceHeader` **must** be a valid pointer to a valid `StdVideoAV1SequenceHeader` value

## [](#_see_also)See Also

[VK\_KHR\_video\_decode\_av1](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_video_decode_av1.html), [VK\_KHR\_video\_maintenance2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_video_maintenance2.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkVideoDecodeAV1InlineSessionParametersInfoKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0