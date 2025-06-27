# VkDeviceGroupDeviceCreateInfo(3) Manual Page

## Name

VkDeviceGroupDeviceCreateInfo - Create a logical device from multiple
physical devices



## <a href="#_c_specification" class="anchor"></a>C Specification

A logical device **can** be created that connects to one or more
physical devices by adding a `VkDeviceGroupDeviceCreateInfo` structure
to the `pNext` chain of [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html).
The `VkDeviceGroupDeviceCreateInfo` structure is defined as:

``` c
// Provided by VK_VERSION_1_1
typedef struct VkDeviceGroupDeviceCreateInfo {
    VkStructureType            sType;
    const void*                pNext;
    uint32_t                   physicalDeviceCount;
    const VkPhysicalDevice*    pPhysicalDevices;
} VkDeviceGroupDeviceCreateInfo;
```

or the equivalent

``` c
// Provided by VK_KHR_device_group_creation
typedef VkDeviceGroupDeviceCreateInfo VkDeviceGroupDeviceCreateInfoKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `physicalDeviceCount` is the number of elements in the
  `pPhysicalDevices` array.

- `pPhysicalDevices` is a pointer to an array of physical device handles
  belonging to the same device group.

## <a href="#_description" class="anchor"></a>Description

The elements of the `pPhysicalDevices` array are an ordered list of the
physical devices that the logical device represents. These **must** be a
subset of a single device group, and need not be in the same order as
they were enumerated. The order of the physical devices in the
`pPhysicalDevices` array determines the *device index* of each physical
device, with element i being assigned a device index of i. Certain
commands and structures refer to one or more physical devices by using
device indices or *device masks* formed using device indices.

A logical device created without using `VkDeviceGroupDeviceCreateInfo`,
or with `physicalDeviceCount` equal to zero, is equivalent to a
`physicalDeviceCount` of one and `pPhysicalDevices` pointing to the
`physicalDevice` parameter to [vkCreateDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateDevice.html). In
particular, the device index of that physical device is zero.

Valid Usage

- <a href="#VUID-VkDeviceGroupDeviceCreateInfo-pPhysicalDevices-00375"
  id="VUID-VkDeviceGroupDeviceCreateInfo-pPhysicalDevices-00375"></a>
  VUID-VkDeviceGroupDeviceCreateInfo-pPhysicalDevices-00375  
  Each element of `pPhysicalDevices` **must** be unique

- <a href="#VUID-VkDeviceGroupDeviceCreateInfo-pPhysicalDevices-00376"
  id="VUID-VkDeviceGroupDeviceCreateInfo-pPhysicalDevices-00376"></a>
  VUID-VkDeviceGroupDeviceCreateInfo-pPhysicalDevices-00376  
  All elements of `pPhysicalDevices` **must** be in the same device
  group as enumerated by
  [vkEnumeratePhysicalDeviceGroups](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkEnumeratePhysicalDeviceGroups.html)

- <a href="#VUID-VkDeviceGroupDeviceCreateInfo-physicalDeviceCount-00377"
  id="VUID-VkDeviceGroupDeviceCreateInfo-physicalDeviceCount-00377"></a>
  VUID-VkDeviceGroupDeviceCreateInfo-physicalDeviceCount-00377  
  If `physicalDeviceCount` is not `0`, the `physicalDevice` parameter of
  [vkCreateDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateDevice.html) **must** be an element of
  `pPhysicalDevices`

Valid Usage (Implicit)

- <a href="#VUID-VkDeviceGroupDeviceCreateInfo-sType-sType"
  id="VUID-VkDeviceGroupDeviceCreateInfo-sType-sType"></a>
  VUID-VkDeviceGroupDeviceCreateInfo-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_DEVICE_GROUP_DEVICE_CREATE_INFO`

- <a href="#VUID-VkDeviceGroupDeviceCreateInfo-pPhysicalDevices-parameter"
  id="VUID-VkDeviceGroupDeviceCreateInfo-pPhysicalDevices-parameter"></a>
  VUID-VkDeviceGroupDeviceCreateInfo-pPhysicalDevices-parameter  
  If `physicalDeviceCount` is not `0`, `pPhysicalDevices` **must** be a
  valid pointer to an array of `physicalDeviceCount` valid
  [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevice.html) handles

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_1](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_1.html),
[VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevice.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkDeviceGroupDeviceCreateInfo"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
