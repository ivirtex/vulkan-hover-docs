# VkPhysicalDeviceIndexTypeUint8Features(3) Manual Page

## Name

VkPhysicalDeviceIndexTypeUint8Features - Structure describing whether uint8 index type can be used



## [](#_c_specification)C Specification

The `VkPhysicalDeviceIndexTypeUint8Features` structure is defined as:

```c++
// Provided by VK_VERSION_1_4
typedef struct VkPhysicalDeviceIndexTypeUint8Features {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           indexTypeUint8;
} VkPhysicalDeviceIndexTypeUint8Features;
```

or the equivalent

```c++
// Provided by VK_KHR_index_type_uint8
typedef VkPhysicalDeviceIndexTypeUint8Features VkPhysicalDeviceIndexTypeUint8FeaturesKHR;
```

or the equivalent

```c++
// Provided by VK_EXT_index_type_uint8
typedef VkPhysicalDeviceIndexTypeUint8Features VkPhysicalDeviceIndexTypeUint8FeaturesEXT;
```

## [](#_members)Members

This structure describes the following feature:

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.

## [](#_description)Description

- []()`indexTypeUint8` indicates that `VK_INDEX_TYPE_UINT8` can be used with [vkCmdBindIndexBuffer2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBindIndexBuffer2.html) and [vkCmdBindIndexBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBindIndexBuffer.html).

If the `VkPhysicalDeviceIndexTypeUint8Features` structure is included in the `pNext` chain of the [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html) structure passed to [vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceFeatures2.html), it is filled in to indicate whether each corresponding feature is supported. If the application wishes to use a [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) with any features described by `VkPhysicalDeviceIndexTypeUint8Features`, it **must** add an instance of the structure, with the desired feature members set to `VK_TRUE`, to the `pNext` chain of [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html) when creating the [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html).

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDeviceIndexTypeUint8Features-sType-sType)VUID-VkPhysicalDeviceIndexTypeUint8Features-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_INDEX_TYPE_UINT8_FEATURES`

## [](#_see_also)See Also

[VK\_EXT\_index\_type\_uint8](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_index_type_uint8.html), [VK\_KHR\_index\_type\_uint8](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_index_type_uint8.html), [VK\_VERSION\_1\_4](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_4.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceIndexTypeUint8Features)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0