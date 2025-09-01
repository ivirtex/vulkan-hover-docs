# VkComponentMapping(3) Manual Page

## Name

VkComponentMapping - Structure specifying a color component mapping



## [](#_c_specification)C Specification

The `VkComponentMapping` structure is defined as:

```c++
// Provided by VK_VERSION_1_0
typedef struct VkComponentMapping {
    VkComponentSwizzle    r;
    VkComponentSwizzle    g;
    VkComponentSwizzle    b;
    VkComponentSwizzle    a;
} VkComponentMapping;
```

## [](#_members)Members

- `r` is a [VkComponentSwizzle](https://registry.khronos.org/vulkan/specs/latest/man/html/VkComponentSwizzle.html) specifying the component value placed in the R component of the output vector.
- `g` is a [VkComponentSwizzle](https://registry.khronos.org/vulkan/specs/latest/man/html/VkComponentSwizzle.html) specifying the component value placed in the G component of the output vector.
- `b` is a [VkComponentSwizzle](https://registry.khronos.org/vulkan/specs/latest/man/html/VkComponentSwizzle.html) specifying the component value placed in the B component of the output vector.
- `a` is a [VkComponentSwizzle](https://registry.khronos.org/vulkan/specs/latest/man/html/VkComponentSwizzle.html) specifying the component value placed in the A component of the output vector.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkComponentMapping-r-parameter)VUID-VkComponentMapping-r-parameter  
  `r` **must** be a valid [VkComponentSwizzle](https://registry.khronos.org/vulkan/specs/latest/man/html/VkComponentSwizzle.html) value
- [](#VUID-VkComponentMapping-g-parameter)VUID-VkComponentMapping-g-parameter  
  `g` **must** be a valid [VkComponentSwizzle](https://registry.khronos.org/vulkan/specs/latest/man/html/VkComponentSwizzle.html) value
- [](#VUID-VkComponentMapping-b-parameter)VUID-VkComponentMapping-b-parameter  
  `b` **must** be a valid [VkComponentSwizzle](https://registry.khronos.org/vulkan/specs/latest/man/html/VkComponentSwizzle.html) value
- [](#VUID-VkComponentMapping-a-parameter)VUID-VkComponentMapping-a-parameter  
  `a` **must** be a valid [VkComponentSwizzle](https://registry.khronos.org/vulkan/specs/latest/man/html/VkComponentSwizzle.html) value

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkAndroidHardwareBufferFormatProperties2ANDROID](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAndroidHardwareBufferFormatProperties2ANDROID.html), [VkAndroidHardwareBufferFormatPropertiesANDROID](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAndroidHardwareBufferFormatPropertiesANDROID.html), [VkBufferCollectionPropertiesFUCHSIA](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferCollectionPropertiesFUCHSIA.html), [VkComponentSwizzle](https://registry.khronos.org/vulkan/specs/latest/man/html/VkComponentSwizzle.html), [VkImageViewCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageViewCreateInfo.html), [VkSamplerBorderColorComponentMappingCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerBorderColorComponentMappingCreateInfoEXT.html), [VkSamplerYcbcrConversionCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerYcbcrConversionCreateInfo.html), [VkScreenBufferFormatPropertiesQNX](https://registry.khronos.org/vulkan/specs/latest/man/html/VkScreenBufferFormatPropertiesQNX.html), [VkVideoFormatPropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoFormatPropertiesKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkComponentMapping).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0