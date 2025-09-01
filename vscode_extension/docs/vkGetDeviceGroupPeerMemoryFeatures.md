# vkGetDeviceGroupPeerMemoryFeatures(3) Manual Page

## Name

vkGetDeviceGroupPeerMemoryFeatures - Query supported peer memory features of a device



## [](#_c_specification)C Specification

*Peer memory* is memory that is allocated for a given physical device and then bound to a resource and accessed by a different physical device, in a logical device that represents multiple physical devices. Some ways of reading and writing peer memory **may** not be supported by a device.

To determine how peer memory **can** be accessed, call:

```c++
// Provided by VK_VERSION_1_1
void vkGetDeviceGroupPeerMemoryFeatures(
    VkDevice                                    device,
    uint32_t                                    heapIndex,
    uint32_t                                    localDeviceIndex,
    uint32_t                                    remoteDeviceIndex,
    VkPeerMemoryFeatureFlags*                   pPeerMemoryFeatures);
```

or the equivalent command

```c++
// Provided by VK_KHR_device_group
void vkGetDeviceGroupPeerMemoryFeaturesKHR(
    VkDevice                                    device,
    uint32_t                                    heapIndex,
    uint32_t                                    localDeviceIndex,
    uint32_t                                    remoteDeviceIndex,
    VkPeerMemoryFeatureFlags*                   pPeerMemoryFeatures);
```

## [](#_parameters)Parameters

- `device` is the logical device that owns the memory.
- `heapIndex` is the index of the memory heap from which the memory is allocated.
- `localDeviceIndex` is the device index of the physical device that performs the memory access.
- `remoteDeviceIndex` is the device index of the physical device that the memory is allocated for.
- `pPeerMemoryFeatures` is a pointer to a [VkPeerMemoryFeatureFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPeerMemoryFeatureFlags.html) bitmask indicating which types of memory accesses are supported for the combination of heap, local, and remote devices.

## [](#_description)Description

Valid Usage

- [](#VUID-vkGetDeviceGroupPeerMemoryFeatures-heapIndex-00691)VUID-vkGetDeviceGroupPeerMemoryFeatures-heapIndex-00691  
  `heapIndex` **must** be less than `memoryHeapCount`
- [](#VUID-vkGetDeviceGroupPeerMemoryFeatures-localDeviceIndex-00692)VUID-vkGetDeviceGroupPeerMemoryFeatures-localDeviceIndex-00692  
  `localDeviceIndex` **must** be a valid device index
- [](#VUID-vkGetDeviceGroupPeerMemoryFeatures-remoteDeviceIndex-00693)VUID-vkGetDeviceGroupPeerMemoryFeatures-remoteDeviceIndex-00693  
  `remoteDeviceIndex` **must** be a valid device index
- [](#VUID-vkGetDeviceGroupPeerMemoryFeatures-localDeviceIndex-00694)VUID-vkGetDeviceGroupPeerMemoryFeatures-localDeviceIndex-00694  
  `localDeviceIndex` **must** not equal `remoteDeviceIndex`

Valid Usage (Implicit)

- [](#VUID-vkGetDeviceGroupPeerMemoryFeatures-device-parameter)VUID-vkGetDeviceGroupPeerMemoryFeatures-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkGetDeviceGroupPeerMemoryFeatures-pPeerMemoryFeatures-parameter)VUID-vkGetDeviceGroupPeerMemoryFeatures-pPeerMemoryFeatures-parameter  
  `pPeerMemoryFeatures` **must** be a valid pointer to a [VkPeerMemoryFeatureFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPeerMemoryFeatureFlags.html) value

## [](#_see_also)See Also

[VK\_VERSION\_1\_1](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_1.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkPeerMemoryFeatureFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPeerMemoryFeatureFlags.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkGetDeviceGroupPeerMemoryFeatures).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0