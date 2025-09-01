# VkDeviceGroupDeviceCreateInfo(3) Manual Page

## Name

VkDeviceGroupDeviceCreateInfo - Create a logical device from multiple physical devices



## [](#_c_specification)C Specification

A logical device **can** be created that connects to one or more physical devices by adding a `VkDeviceGroupDeviceCreateInfo` structure to the `pNext` chain of [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html). The `VkDeviceGroupDeviceCreateInfo` structure is defined as:

```c++
// Provided by VK_VERSION_1_1
typedef struct VkDeviceGroupDeviceCreateInfo {
    VkStructureType            sType;
    const void*                pNext;
    uint32_t                   physicalDeviceCount;
    const VkPhysicalDevice*    pPhysicalDevices;
} VkDeviceGroupDeviceCreateInfo;
```

or the equivalent

```c++
// Provided by VK_KHR_device_group_creation
typedef VkDeviceGroupDeviceCreateInfo VkDeviceGroupDeviceCreateInfoKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `physicalDeviceCount` is the number of elements in the `pPhysicalDevices` array.
- `pPhysicalDevices` is a pointer to an array of physical device handles belonging to the same device group.

## [](#_description)Description

The elements of the `pPhysicalDevices` array are an ordered list of the physical devices that the logical device represents. These **must** be a subset of a single device group, and need not be in the same order as they were enumerated. The order of the physical devices in the `pPhysicalDevices` array determines the *device index* of each physical device, with element i being assigned a device index of i. Certain commands and structures refer to one or more physical devices by using device indices or *device masks* formed using device indices.

A logical device created without using `VkDeviceGroupDeviceCreateInfo`, or with `physicalDeviceCount` equal to zero, is equivalent to a `physicalDeviceCount` of one and `pPhysicalDevices` pointing to the `physicalDevice` parameter to [vkCreateDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateDevice.html). In particular, the device index of that physical device is zero.

Valid Usage

- [](#VUID-VkDeviceGroupDeviceCreateInfo-pPhysicalDevices-00375)VUID-VkDeviceGroupDeviceCreateInfo-pPhysicalDevices-00375  
  Each element of `pPhysicalDevices` **must** be unique
- [](#VUID-VkDeviceGroupDeviceCreateInfo-pPhysicalDevices-00376)VUID-VkDeviceGroupDeviceCreateInfo-pPhysicalDevices-00376  
  All elements of `pPhysicalDevices` **must** be in the same device group as enumerated by [vkEnumeratePhysicalDeviceGroups](https://registry.khronos.org/vulkan/specs/latest/man/html/vkEnumeratePhysicalDeviceGroups.html)
- [](#VUID-VkDeviceGroupDeviceCreateInfo-physicalDeviceCount-00377)VUID-VkDeviceGroupDeviceCreateInfo-physicalDeviceCount-00377  
  If `physicalDeviceCount` is not `0`, the `physicalDevice` parameter of [vkCreateDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateDevice.html) **must** be an element of `pPhysicalDevices`

Valid Usage (Implicit)

- [](#VUID-VkDeviceGroupDeviceCreateInfo-sType-sType)VUID-VkDeviceGroupDeviceCreateInfo-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_DEVICE_GROUP_DEVICE_CREATE_INFO`
- [](#VUID-VkDeviceGroupDeviceCreateInfo-pPhysicalDevices-parameter)VUID-VkDeviceGroupDeviceCreateInfo-pPhysicalDevices-parameter  
  If `physicalDeviceCount` is not `0`, `pPhysicalDevices` **must** be a valid pointer to an array of `physicalDeviceCount` valid [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevice.html) handles

## [](#_see_also)See Also

[VK\_VERSION\_1\_1](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_1.html), [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevice.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkDeviceGroupDeviceCreateInfo).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0