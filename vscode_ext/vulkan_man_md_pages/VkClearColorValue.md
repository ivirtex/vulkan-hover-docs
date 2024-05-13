# VkClearColorValue(3) Manual Page

## Name

VkClearColorValue - Structure specifying a clear color value



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkClearColorValue` structure is defined as:

``` c
// Provided by VK_VERSION_1_0
typedef union VkClearColorValue {
    float       float32[4];
    int32_t     int32[4];
    uint32_t    uint32[4];
} VkClearColorValue;
```

## <a href="#_members" class="anchor"></a>Members

- `float32` are the color clear values when the format of the image or
  attachment is one of the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#formats-numericformat"
  target="_blank" rel="noopener">numeric formats</a> with a numeric type
  that is floating-point. Floating point values are automatically
  converted to the format of the image, with the clear value being
  treated as linear if the image is sRGB.

- `int32` are the color clear values when the format of the image or
  attachment has a numeric type that is signed integer (`SINT`). Signed
  integer values are converted to the format of the image by casting to
  the smaller type (with negative 32-bit values mapping to negative
  values in the smaller type). If the integer clear value is not
  representable in the target type (e.g. would overflow in conversion to
  that type), the clear value is undefined.

- `uint32` are the color clear values when the format of the image or
  attachment has a numeric type that is unsigned integer (`UINT`).
  Unsigned integer values are converted to the format of the image by
  casting to the integer type with fewer bits.

## <a href="#_description" class="anchor"></a>Description

The four array elements of the clear color map to R, G, B, and A
components of image formats, in order.

If the image has more than one sample, the same value is written to all
samples for any pixels being cleared.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkClearValue](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkClearValue.html),
[VkSamplerCustomBorderColorCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerCustomBorderColorCreateInfoEXT.html),
[vkCmdClearColorImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdClearColorImage.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkClearColorValue"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
