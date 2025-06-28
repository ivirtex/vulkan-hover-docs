# VkPhysicalDeviceSchedulingControlsPropertiesARM(3) Manual Page

## Name

VkPhysicalDeviceSchedulingControlsPropertiesARM - Structure containing scheduling control properties of a physical device



## [](#_c_specification)C Specification

The `VkPhysicalDeviceSchedulingControlsPropertiesARM` structure is defined as:

```c++
// Provided by VK_ARM_scheduling_controls
typedef struct VkPhysicalDeviceSchedulingControlsPropertiesARM {
    VkStructureType                               sType;
    void*                                         pNext;
    VkPhysicalDeviceSchedulingControlsFlagsARM    schedulingControlsFlags;
} VkPhysicalDeviceSchedulingControlsPropertiesARM;
```

## [](#_members)Members

- []()`schedulingControlsFlags` specifies the specific scheduling controls that a physical device supports.

## [](#_description)Description

If the `VkPhysicalDeviceSchedulingControlsPropertiesARM` structure is included in the `pNext` chain of the [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html) structure passed to [vkGetPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceProperties2.html), it is filled in with each corresponding implementation-dependent property.

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDeviceSchedulingControlsPropertiesARM-sType-sType)VUID-VkPhysicalDeviceSchedulingControlsPropertiesARM-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SCHEDULING_CONTROLS_PROPERTIES_ARM`

## [](#_see_also)See Also

[VK\_ARM\_scheduling\_controls](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_ARM_scheduling_controls.html), [VkPhysicalDeviceSchedulingControlsFlagsARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceSchedulingControlsFlagsARM.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceSchedulingControlsPropertiesARM)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0