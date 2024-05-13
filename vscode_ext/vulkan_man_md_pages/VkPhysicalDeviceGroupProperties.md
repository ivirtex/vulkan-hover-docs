# VkPhysicalDeviceGroupProperties(3) Manual Page

## Name

VkPhysicalDeviceGroupProperties - Structure specifying physical device
group properties



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceGroupProperties` structure is defined as:

``` c
// Provided by VK_VERSION_1_1
typedef struct VkPhysicalDeviceGroupProperties {
    VkStructureType     sType;
    void*               pNext;
    uint32_t            physicalDeviceCount;
    VkPhysicalDevice    physicalDevices[VK_MAX_DEVICE_GROUP_SIZE];
    VkBool32            subsetAllocation;
} VkPhysicalDeviceGroupProperties;
```

or the equivalent

``` c
// Provided by VK_KHR_device_group_creation
typedef VkPhysicalDeviceGroupProperties VkPhysicalDeviceGroupPropertiesKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `physicalDeviceCount` is the number of physical devices in the group.

- `physicalDevices` is an array of `VK_MAX_DEVICE_GROUP_SIZE`
  [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevice.html) handles representing all
  physical devices in the group. The first `physicalDeviceCount`
  elements of the array will be valid.

- `subsetAllocation` specifies whether logical devices created from the
  group support allocating device memory on a subset of devices, via the
  `deviceMask` member of the
  [VkMemoryAllocateFlagsInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryAllocateFlagsInfo.html). If this
  is `VK_FALSE`, then all device memory allocations are made across all
  physical devices in the group. If `physicalDeviceCount` is `1`, then
  `subsetAllocation` **must** be `VK_FALSE`.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-VkPhysicalDeviceGroupProperties-sType-sType"
  id="VUID-VkPhysicalDeviceGroupProperties-sType-sType"></a>
  VUID-VkPhysicalDeviceGroupProperties-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_GROUP_PROPERTIES`

- <a href="#VUID-VkPhysicalDeviceGroupProperties-pNext-pNext"
  id="VUID-VkPhysicalDeviceGroupProperties-pNext-pNext"></a>
  VUID-VkPhysicalDeviceGroupProperties-pNext-pNext  
  `pNext` **must** be `NULL`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_1](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_1.html), [VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html),
[VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevice.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkEnumeratePhysicalDeviceGroups](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkEnumeratePhysicalDeviceGroups.html),
[vkEnumeratePhysicalDeviceGroupsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkEnumeratePhysicalDeviceGroupsKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceGroupProperties"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
