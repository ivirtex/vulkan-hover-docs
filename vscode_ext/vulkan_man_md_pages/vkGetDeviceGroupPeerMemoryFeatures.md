# vkGetDeviceGroupPeerMemoryFeatures(3) Manual Page

## Name

vkGetDeviceGroupPeerMemoryFeatures - Query supported peer memory
features of a device



## <a href="#_c_specification" class="anchor"></a>C Specification

*Peer memory* is memory that is allocated for a given physical device
and then bound to a resource and accessed by a different physical
device, in a logical device that represents multiple physical devices.
Some ways of reading and writing peer memory **may** not be supported by
a device.

To determine how peer memory **can** be accessed, call:

``` c
// Provided by VK_VERSION_1_1
void vkGetDeviceGroupPeerMemoryFeatures(
    VkDevice                                    device,
    uint32_t                                    heapIndex,
    uint32_t                                    localDeviceIndex,
    uint32_t                                    remoteDeviceIndex,
    VkPeerMemoryFeatureFlags*                   pPeerMemoryFeatures);
```

or the equivalent command

``` c
// Provided by VK_KHR_device_group
void vkGetDeviceGroupPeerMemoryFeaturesKHR(
    VkDevice                                    device,
    uint32_t                                    heapIndex,
    uint32_t                                    localDeviceIndex,
    uint32_t                                    remoteDeviceIndex,
    VkPeerMemoryFeatureFlags*                   pPeerMemoryFeatures);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device that owns the memory.

- `heapIndex` is the index of the memory heap from which the memory is
  allocated.

- `localDeviceIndex` is the device index of the physical device that
  performs the memory access.

- `remoteDeviceIndex` is the device index of the physical device that
  the memory is allocated for.

- `pPeerMemoryFeatures` is a pointer to a
  [VkPeerMemoryFeatureFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPeerMemoryFeatureFlags.html) bitmask
  indicating which types of memory accesses are supported for the
  combination of heap, local, and remote devices.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-vkGetDeviceGroupPeerMemoryFeatures-heapIndex-00691"
  id="VUID-vkGetDeviceGroupPeerMemoryFeatures-heapIndex-00691"></a>
  VUID-vkGetDeviceGroupPeerMemoryFeatures-heapIndex-00691  
  `heapIndex` **must** be less than `memoryHeapCount`

- <a
  href="#VUID-vkGetDeviceGroupPeerMemoryFeatures-localDeviceIndex-00692"
  id="VUID-vkGetDeviceGroupPeerMemoryFeatures-localDeviceIndex-00692"></a>
  VUID-vkGetDeviceGroupPeerMemoryFeatures-localDeviceIndex-00692  
  `localDeviceIndex` **must** be a valid device index

- <a
  href="#VUID-vkGetDeviceGroupPeerMemoryFeatures-remoteDeviceIndex-00693"
  id="VUID-vkGetDeviceGroupPeerMemoryFeatures-remoteDeviceIndex-00693"></a>
  VUID-vkGetDeviceGroupPeerMemoryFeatures-remoteDeviceIndex-00693  
  `remoteDeviceIndex` **must** be a valid device index

- <a
  href="#VUID-vkGetDeviceGroupPeerMemoryFeatures-localDeviceIndex-00694"
  id="VUID-vkGetDeviceGroupPeerMemoryFeatures-localDeviceIndex-00694"></a>
  VUID-vkGetDeviceGroupPeerMemoryFeatures-localDeviceIndex-00694  
  `localDeviceIndex` **must** not equal `remoteDeviceIndex`

Valid Usage (Implicit)

- <a href="#VUID-vkGetDeviceGroupPeerMemoryFeatures-device-parameter"
  id="VUID-vkGetDeviceGroupPeerMemoryFeatures-device-parameter"></a>
  VUID-vkGetDeviceGroupPeerMemoryFeatures-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a
  href="#VUID-vkGetDeviceGroupPeerMemoryFeatures-pPeerMemoryFeatures-parameter"
  id="VUID-vkGetDeviceGroupPeerMemoryFeatures-pPeerMemoryFeatures-parameter"></a>
  VUID-vkGetDeviceGroupPeerMemoryFeatures-pPeerMemoryFeatures-parameter  
  `pPeerMemoryFeatures` **must** be a valid pointer to a
  [VkPeerMemoryFeatureFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPeerMemoryFeatureFlags.html) value

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_1](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_1.html), [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html),
[VkPeerMemoryFeatureFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPeerMemoryFeatureFlags.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkGetDeviceGroupPeerMemoryFeatures"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
