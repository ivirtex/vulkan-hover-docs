# VkImageViewASTCDecodeModeEXT(3) Manual Page

## Name

VkImageViewASTCDecodeModeEXT - Structure describing the ASTC decode mode for an image view



## [](#_c_specification)C Specification

If the `pNext` chain includes a `VkImageViewASTCDecodeModeEXT` structure, then that structure includes a parameter specifying the decode mode for image views using ASTC compressed formats.

The `VkImageViewASTCDecodeModeEXT` structure is defined as:

```c++
// Provided by VK_EXT_astc_decode_mode
typedef struct VkImageViewASTCDecodeModeEXT {
    VkStructureType    sType;
    const void*        pNext;
    VkFormat           decodeMode;
} VkImageViewASTCDecodeModeEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `decodeMode` is the intermediate format used to decode ASTC compressed formats.

## [](#_description)Description

Valid Usage

- [](#VUID-VkImageViewASTCDecodeModeEXT-decodeMode-02230)VUID-VkImageViewASTCDecodeModeEXT-decodeMode-02230  
  `decodeMode` **must** be one of `VK_FORMAT_R16G16B16A16_SFLOAT`, `VK_FORMAT_R8G8B8A8_UNORM`, or `VK_FORMAT_E5B9G9R9_UFLOAT_PACK32`
- [](#VUID-VkImageViewASTCDecodeModeEXT-decodeMode-02231)VUID-VkImageViewASTCDecodeModeEXT-decodeMode-02231  
  If the [`decodeModeSharedExponent`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-astc-decodeModeSharedExponent) feature is not enabled, `decodeMode` **must** not be `VK_FORMAT_E5B9G9R9_UFLOAT_PACK32`
- [](#VUID-VkImageViewASTCDecodeModeEXT-decodeMode-02232)VUID-VkImageViewASTCDecodeModeEXT-decodeMode-02232  
  If `decodeMode` is `VK_FORMAT_R8G8B8A8_UNORM` the image view **must** not include blocks using any of the ASTC HDR modes
- [](#VUID-VkImageViewASTCDecodeModeEXT-format-04084)VUID-VkImageViewASTCDecodeModeEXT-format-04084  
  `format` of the image view **must** be one of the [ASTC Compressed Image Formats](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#appendix-compressedtex-astc)

If `format` uses sRGB encoding then the `decodeMode` has no effect.

Valid Usage (Implicit)

- [](#VUID-VkImageViewASTCDecodeModeEXT-sType-sType)VUID-VkImageViewASTCDecodeModeEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_IMAGE_VIEW_ASTC_DECODE_MODE_EXT`
- [](#VUID-VkImageViewASTCDecodeModeEXT-decodeMode-parameter)VUID-VkImageViewASTCDecodeModeEXT-decodeMode-parameter  
  `decodeMode` **must** be a valid [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) value

## [](#_see_also)See Also

[VK\_EXT\_astc\_decode\_mode](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_astc_decode_mode.html), [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkImageViewASTCDecodeModeEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0