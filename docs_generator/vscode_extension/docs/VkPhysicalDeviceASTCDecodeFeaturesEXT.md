# VkPhysicalDeviceASTCDecodeFeaturesEXT(3) Manual Page

## Name

VkPhysicalDeviceASTCDecodeFeaturesEXT - Structure describing ASTC decode mode features



## [](#_c_specification)C Specification

The `VkPhysicalDeviceASTCDecodeFeaturesEXT` structure is defined as:

```c++
// Provided by VK_EXT_astc_decode_mode
typedef struct VkPhysicalDeviceASTCDecodeFeaturesEXT {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           decodeModeSharedExponent;
} VkPhysicalDeviceASTCDecodeFeaturesEXT;
```

## [](#_members)Members

This structure describes the following feature:

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- []()`decodeModeSharedExponent` indicates whether the implementation supports decoding ASTC compressed formats to `VK_FORMAT_E5B9G9R9_UFLOAT_PACK32` internal precision.

## [](#_description)Description

If the `VkPhysicalDeviceASTCDecodeFeaturesEXT` structure is included in the `pNext` chain of the [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html) structure passed to [vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceFeatures2.html), it is filled in to indicate whether each corresponding feature is supported. If the application wishes to use a [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) with any features described by `VkPhysicalDeviceASTCDecodeFeaturesEXT`, it **must** add an instance of the structure, with the desired feature members set to `VK_TRUE`, to the `pNext` chain of [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html) when creating the [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html).

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDeviceASTCDecodeFeaturesEXT-sType-sType)VUID-VkPhysicalDeviceASTCDecodeFeaturesEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_ASTC_DECODE_FEATURES_EXT`

## [](#_see_also)See Also

[VK\_EXT\_astc\_decode\_mode](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_astc_decode_mode.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceASTCDecodeFeaturesEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0