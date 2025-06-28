# VkDeviceGroupPresentCapabilitiesKHR(3) Manual Page

## Name

VkDeviceGroupPresentCapabilitiesKHR - Present capabilities from other physical devices



## [](#_c_specification)C Specification

The `VkDeviceGroupPresentCapabilitiesKHR` structure is defined as:

```c++
// Provided by VK_VERSION_1_1 with VK_KHR_swapchain, VK_KHR_device_group with VK_KHR_surface
typedef struct VkDeviceGroupPresentCapabilitiesKHR {
    VkStructureType                     sType;
    void*                               pNext;
    uint32_t                            presentMask[VK_MAX_DEVICE_GROUP_SIZE];
    VkDeviceGroupPresentModeFlagsKHR    modes;
} VkDeviceGroupPresentCapabilitiesKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `presentMask` is an array of `VK_MAX_DEVICE_GROUP_SIZE` `uint32_t` masks, where the mask at element i is non-zero if physical device i has a presentation engine, and where bit j is set in element i if physical device i **can** present swapchain images from physical device j. If element i is non-zero, then bit i **must** be set.
- `modes` is a bitmask of [VkDeviceGroupPresentModeFlagBitsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceGroupPresentModeFlagBitsKHR.html) indicating which device group presentation modes are supported.

## [](#_description)Description

`modes` always has `VK_DEVICE_GROUP_PRESENT_MODE_LOCAL_BIT_KHR` set.

The present mode flags are also used when presenting an image, in [VkDeviceGroupPresentInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceGroupPresentInfoKHR.html)::`mode`.

If a device group only includes a single physical device, then `modes` **must** equal `VK_DEVICE_GROUP_PRESENT_MODE_LOCAL_BIT_KHR`.

Valid Usage (Implicit)

- [](#VUID-VkDeviceGroupPresentCapabilitiesKHR-sType-sType)VUID-VkDeviceGroupPresentCapabilitiesKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_DEVICE_GROUP_PRESENT_CAPABILITIES_KHR`
- [](#VUID-VkDeviceGroupPresentCapabilitiesKHR-pNext-pNext)VUID-VkDeviceGroupPresentCapabilitiesKHR-pNext-pNext  
  `pNext` **must** be `NULL`

## [](#_see_also)See Also

[VK\_KHR\_device\_group](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_device_group.html), [VK\_KHR\_surface](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_surface.html), [VK\_KHR\_swapchain](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_swapchain.html), [VK\_VERSION\_1\_1](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_1.html), [VkDeviceGroupPresentModeFlagsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceGroupPresentModeFlagsKHR.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkGetDeviceGroupPresentCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetDeviceGroupPresentCapabilitiesKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkDeviceGroupPresentCapabilitiesKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0