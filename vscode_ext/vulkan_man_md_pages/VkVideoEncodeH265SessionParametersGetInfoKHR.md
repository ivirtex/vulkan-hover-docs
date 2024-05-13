# VkVideoEncodeH265SessionParametersGetInfoKHR(3) Manual Page

## Name

VkVideoEncodeH265SessionParametersGetInfoKHR - Structure specifying
parameters for retrieving encoded H.265 parameter set data



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkVideoEncodeH265SessionParametersGetInfoKHR` structure is defined
as:

``` c
// Provided by VK_KHR_video_encode_h265
typedef struct VkVideoEncodeH265SessionParametersGetInfoKHR {
    VkStructureType    sType;
    const void*        pNext;
    VkBool32           writeStdVPS;
    VkBool32           writeStdSPS;
    VkBool32           writeStdPPS;
    uint32_t           stdVPSId;
    uint32_t           stdSPSId;
    uint32_t           stdPPSId;
} VkVideoEncodeH265SessionParametersGetInfoKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `writeStdVPS` indicates whether the encoded <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-h265-vps"
  target="_blank" rel="noopener">H.265 video parameter set</a>
  identified by `stdVPSId` is requested to be retrieved.

- `writeStdSPS` indicates whether the encoded <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-h265-sps"
  target="_blank" rel="noopener">H.265 sequence parameter set</a>
  identified by the pair constructed from `stdVPSId` and `stdSPSId` is
  requested to be retrieved.

- `writeStdPPS` indicates whether the encoded <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-h265-pps"
  target="_blank" rel="noopener">H.265 picture parameter set</a>
  identified by the triplet constructed from `stdVPSId`, `stdSPSId`, and
  `stdPPSId` is requested to be retrieved.

- `stdVPSId` specifies the H.265 video parameter set ID used to identify
  the retrieved H.265 video, sequence, and/or picture parameter set(s).

- `stdSPSId` specifies the H.265 sequence parameter set ID used to
  identify the retrieved H.265 sequence and/or picture parameter set(s)
  when `writeStdSPS` and/or `writeStdPPS` is set to `VK_TRUE`.

- `stdPPSId` specifies the H.265 picture parameter set ID used to
  identify the retrieved H.265 picture parameter set when `writeStdPPS`
  is set to `VK_TRUE`.

## <a href="#_description" class="anchor"></a>Description

When this structure is specified in the `pNext` chain of the
[VkVideoEncodeSessionParametersGetInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeSessionParametersGetInfoKHR.html)
structure passed to
[vkGetEncodedVideoSessionParametersKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetEncodedVideoSessionParametersKHR.html),
the command will write encoded parameter data to the output buffer in
the following order:

1.  The <a
    href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-h265-vps"
    target="_blank" rel="noopener">H.265 video parameter set</a>
    identified by `stdVPSId`, if `writeStdVPS` is set to `VK_TRUE`.

2.  The <a
    href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-h265-sps"
    target="_blank" rel="noopener">H.265 sequence parameter set</a>
    identified by the pair constructed from `stdVPSId` and `stdSPSId`,
    if `writeStdSPS` is set to `VK_TRUE`.

3.  The <a
    href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-h265-pps"
    target="_blank" rel="noopener">H.265 picture parameter set</a>
    identified by the triplet constructed from `stdVPSId`, `stdSPSId`,
    and `stdPPSId`, if `writeStdPPS` is set to `VK_TRUE`.

Valid Usage

- <a
  href="#VUID-VkVideoEncodeH265SessionParametersGetInfoKHR-writeStdVPS-08290"
  id="VUID-VkVideoEncodeH265SessionParametersGetInfoKHR-writeStdVPS-08290"></a>
  VUID-VkVideoEncodeH265SessionParametersGetInfoKHR-writeStdVPS-08290  
  At least one of `writeStdVPS`, `writeStdSPS`, and `writeStdPPS`
  **must** be set to `VK_TRUE`

Valid Usage (Implicit)

- <a href="#VUID-VkVideoEncodeH265SessionParametersGetInfoKHR-sType-sType"
  id="VUID-VkVideoEncodeH265SessionParametersGetInfoKHR-sType-sType"></a>
  VUID-VkVideoEncodeH265SessionParametersGetInfoKHR-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_VIDEO_ENCODE_H265_SESSION_PARAMETERS_GET_INFO_KHR`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_video_encode_h265](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_video_encode_h265.html),
[VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkVideoEncodeH265SessionParametersGetInfoKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
