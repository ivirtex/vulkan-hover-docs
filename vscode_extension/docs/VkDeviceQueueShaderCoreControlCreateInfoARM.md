# VkDeviceQueueShaderCoreControlCreateInfoARM(3) Manual Page

## Name

VkDeviceQueueShaderCoreControlCreateInfoARM - Control the number of shader cores used by queues



## [](#_c_specification)C Specification

The number of shader cores used by a queue **can** be controlled by adding a `VkDeviceQueueShaderCoreControlCreateInfoARM` structure to the `pNext` chain of [VkDeviceQueueCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceQueueCreateInfo.html) structures.

The `VkDeviceQueueShaderCoreControlCreateInfoARM` structure is defined as:

```c++
// Provided by VK_ARM_scheduling_controls
typedef struct VkDeviceQueueShaderCoreControlCreateInfoARM {
    VkStructureType    sType;
    void*              pNext;
    uint32_t           shaderCoreCount;
} VkDeviceQueueShaderCoreControlCreateInfoARM;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `shaderCoreCount` is the number of shader cores this queue uses.

## [](#_description)Description

Queues created without specifying `VkDeviceQueueShaderCoreControlCreateInfoARM` will default to using all the shader cores available.

Valid Usage

- [](#VUID-VkDeviceQueueShaderCoreControlCreateInfoARM-shaderCoreCount-09399)VUID-VkDeviceQueueShaderCoreControlCreateInfoARM-shaderCoreCount-09399  
  `shaderCoreCount` **must** be greater than 0 and less than or equal to the total number of shader cores as reported via [VkPhysicalDeviceShaderCoreBuiltinsPropertiesARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceShaderCoreBuiltinsPropertiesARM.html)::`shaderCoreCount`

Valid Usage (Implicit)

- [](#VUID-VkDeviceQueueShaderCoreControlCreateInfoARM-sType-sType)VUID-VkDeviceQueueShaderCoreControlCreateInfoARM-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_DEVICE_QUEUE_SHADER_CORE_CONTROL_CREATE_INFO_ARM`

## [](#_see_also)See Also

[VK\_ARM\_scheduling\_controls](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_ARM_scheduling_controls.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkDeviceQueueShaderCoreControlCreateInfoARM).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0