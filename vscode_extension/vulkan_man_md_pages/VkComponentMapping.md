# VkComponentMapping(3) Manual Page

## Name

VkComponentMapping - Structure specifying a color component mapping



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkComponentMapping` structure is defined as:

``` c
// Provided by VK_VERSION_1_0
typedef struct VkComponentMapping {
    VkComponentSwizzle    r;
    VkComponentSwizzle    g;
    VkComponentSwizzle    b;
    VkComponentSwizzle    a;
} VkComponentMapping;
```

## <a href="#_members" class="anchor"></a>Members

- `r` is a [VkComponentSwizzle](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkComponentSwizzle.html) specifying the
  component value placed in the R component of the output vector.

- `g` is a [VkComponentSwizzle](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkComponentSwizzle.html) specifying the
  component value placed in the G component of the output vector.

- `b` is a [VkComponentSwizzle](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkComponentSwizzle.html) specifying the
  component value placed in the B component of the output vector.

- `a` is a [VkComponentSwizzle](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkComponentSwizzle.html) specifying the
  component value placed in the A component of the output vector.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-VkComponentMapping-r-parameter"
  id="VUID-VkComponentMapping-r-parameter"></a>
  VUID-VkComponentMapping-r-parameter  
  `r` **must** be a valid [VkComponentSwizzle](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkComponentSwizzle.html)
  value

- <a href="#VUID-VkComponentMapping-g-parameter"
  id="VUID-VkComponentMapping-g-parameter"></a>
  VUID-VkComponentMapping-g-parameter  
  `g` **must** be a valid [VkComponentSwizzle](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkComponentSwizzle.html)
  value

- <a href="#VUID-VkComponentMapping-b-parameter"
  id="VUID-VkComponentMapping-b-parameter"></a>
  VUID-VkComponentMapping-b-parameter  
  `b` **must** be a valid [VkComponentSwizzle](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkComponentSwizzle.html)
  value

- <a href="#VUID-VkComponentMapping-a-parameter"
  id="VUID-VkComponentMapping-a-parameter"></a>
  VUID-VkComponentMapping-a-parameter  
  `a` **must** be a valid [VkComponentSwizzle](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkComponentSwizzle.html)
  value

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkAndroidHardwareBufferFormatProperties2ANDROID](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAndroidHardwareBufferFormatProperties2ANDROID.html),
[VkAndroidHardwareBufferFormatPropertiesANDROID](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAndroidHardwareBufferFormatPropertiesANDROID.html),
[VkBufferCollectionPropertiesFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferCollectionPropertiesFUCHSIA.html),
[VkComponentSwizzle](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkComponentSwizzle.html),
[VkImageViewCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageViewCreateInfo.html),
[VkSamplerBorderColorComponentMappingCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerBorderColorComponentMappingCreateInfoEXT.html),
[VkSamplerYcbcrConversionCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerYcbcrConversionCreateInfo.html),
[VkScreenBufferFormatPropertiesQNX](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkScreenBufferFormatPropertiesQNX.html),
[VkVideoFormatPropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoFormatPropertiesKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkComponentMapping"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
