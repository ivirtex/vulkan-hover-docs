# VkImageViewASTCDecodeModeEXT(3) Manual Page

## Name

VkImageViewASTCDecodeModeEXT - Structure describing the ASTC decode mode
for an image view



## <a href="#_c_specification" class="anchor"></a>C Specification

If the `pNext` chain includes a `VkImageViewASTCDecodeModeEXT`
structure, then that structure includes a parameter specifying the
decode mode for image views using ASTC compressed formats.

The `VkImageViewASTCDecodeModeEXT` structure is defined as:

``` c
// Provided by VK_EXT_astc_decode_mode
typedef struct VkImageViewASTCDecodeModeEXT {
    VkStructureType    sType;
    const void*        pNext;
    VkFormat           decodeMode;
} VkImageViewASTCDecodeModeEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `decodeMode` is the intermediate format used to decode ASTC compressed
  formats.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-VkImageViewASTCDecodeModeEXT-decodeMode-02230"
  id="VUID-VkImageViewASTCDecodeModeEXT-decodeMode-02230"></a>
  VUID-VkImageViewASTCDecodeModeEXT-decodeMode-02230  
  `decodeMode` **must** be one of `VK_FORMAT_R16G16B16A16_SFLOAT`,
  `VK_FORMAT_R8G8B8A8_UNORM`, or `VK_FORMAT_E5B9G9R9_UFLOAT_PACK32`

- <a href="#VUID-VkImageViewASTCDecodeModeEXT-decodeMode-02231"
  id="VUID-VkImageViewASTCDecodeModeEXT-decodeMode-02231"></a>
  VUID-VkImageViewASTCDecodeModeEXT-decodeMode-02231  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-astc-decodeModeSharedExponent"
  target="_blank" rel="noopener"><code>decodeModeSharedExponent</code></a>
  feature is not enabled, `decodeMode` **must** not be
  `VK_FORMAT_E5B9G9R9_UFLOAT_PACK32`

- <a href="#VUID-VkImageViewASTCDecodeModeEXT-decodeMode-02232"
  id="VUID-VkImageViewASTCDecodeModeEXT-decodeMode-02232"></a>
  VUID-VkImageViewASTCDecodeModeEXT-decodeMode-02232  
  If `decodeMode` is `VK_FORMAT_R8G8B8A8_UNORM` the image view **must**
  not include blocks using any of the ASTC HDR modes

- <a href="#VUID-VkImageViewASTCDecodeModeEXT-format-04084"
  id="VUID-VkImageViewASTCDecodeModeEXT-format-04084"></a>
  VUID-VkImageViewASTCDecodeModeEXT-format-04084  
  `format` of the image view **must** be one of the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#appendix-compressedtex-astc"
  target="_blank" rel="noopener">ASTC Compressed Image Formats</a>

If `format` uses sRGB encoding then the `decodeMode` has no effect.

Valid Usage (Implicit)

- <a href="#VUID-VkImageViewASTCDecodeModeEXT-sType-sType"
  id="VUID-VkImageViewASTCDecodeModeEXT-sType-sType"></a>
  VUID-VkImageViewASTCDecodeModeEXT-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_IMAGE_VIEW_ASTC_DECODE_MODE_EXT`

- <a href="#VUID-VkImageViewASTCDecodeModeEXT-decodeMode-parameter"
  id="VUID-VkImageViewASTCDecodeModeEXT-decodeMode-parameter"></a>
  VUID-VkImageViewASTCDecodeModeEXT-decodeMode-parameter  
  `decodeMode` **must** be a valid [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) value

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_astc_decode_mode](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_astc_decode_mode.html),
[VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkImageViewASTCDecodeModeEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
