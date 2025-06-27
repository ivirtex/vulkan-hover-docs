# VkDeviceQueueShaderCoreControlCreateInfoARM(3) Manual Page

## Name

VkDeviceQueueShaderCoreControlCreateInfoARM - Control the number of
shader cores used by queues



## <a href="#_c_specification" class="anchor"></a>C Specification

The number of shader cores used by a queue **can** be controlled by
adding a `VkDeviceQueueShaderCoreControlCreateInfoARM` structure to the
`pNext` chain of [VkDeviceQueueCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceQueueCreateInfo.html)
structures.

The `VkDeviceQueueShaderCoreControlCreateInfoARM` structure is defined
as:

``` c
// Provided by VK_ARM_scheduling_controls
typedef struct VkDeviceQueueShaderCoreControlCreateInfoARM {
    VkStructureType    sType;
    void*              pNext;
    uint32_t           shaderCoreCount;
} VkDeviceQueueShaderCoreControlCreateInfoARM;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `shaderCoreCount` is the number of shader cores this queue uses.

## <a href="#_description" class="anchor"></a>Description

Queues created without specifying
`VkDeviceQueueShaderCoreControlCreateInfoARM` will default to using all
the shader cores available.

Valid Usage

- <a
  href="#VUID-VkDeviceQueueShaderCoreControlCreateInfoARM-shaderCoreCount-09399"
  id="VUID-VkDeviceQueueShaderCoreControlCreateInfoARM-shaderCoreCount-09399"></a>
  VUID-VkDeviceQueueShaderCoreControlCreateInfoARM-shaderCoreCount-09399  
  `shaderCoreCount` **must** be greater than 0 and less than or equal to
  the total number of shader cores as reported via
  [VkPhysicalDeviceShaderCoreBuiltinsPropertiesARM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceShaderCoreBuiltinsPropertiesARM.html)::`shaderCoreCount`

Valid Usage (Implicit)

- <a href="#VUID-VkDeviceQueueShaderCoreControlCreateInfoARM-sType-sType"
  id="VUID-VkDeviceQueueShaderCoreControlCreateInfoARM-sType-sType"></a>
  VUID-VkDeviceQueueShaderCoreControlCreateInfoARM-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_DEVICE_QUEUE_SHADER_CORE_CONTROL_CREATE_INFO_ARM`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_ARM_scheduling_controls](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_ARM_scheduling_controls.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkDeviceQueueShaderCoreControlCreateInfoARM"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
