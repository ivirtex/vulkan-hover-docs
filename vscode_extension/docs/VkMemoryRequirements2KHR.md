# VkMemoryRequirements2(3) Manual Page

## Name

VkMemoryRequirements2 - Structure specifying memory requirements



## [](#_c_specification)C Specification

The `VkMemoryRequirements2` structure is defined as:

```c++
// Provided by VK_VERSION_1_1
typedef struct VkMemoryRequirements2 {
    VkStructureType         sType;
    void*                   pNext;
    VkMemoryRequirements    memoryRequirements;
} VkMemoryRequirements2;
```

or the equivalent

```c++
// Provided by VK_KHR_get_memory_requirements2, VK_NV_ray_tracing with VK_KHR_get_memory_requirements2 or VK_VERSION_1_1
typedef VkMemoryRequirements2 VkMemoryRequirements2KHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `memoryRequirements` is a [VkMemoryRequirements](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryRequirements.html) structure describing the memory requirements of the resource.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkMemoryRequirements2-sType-sType)VUID-VkMemoryRequirements2-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_MEMORY_REQUIREMENTS_2`
- [](#VUID-VkMemoryRequirements2-pNext-pNext)VUID-VkMemoryRequirements2-pNext-pNext  
  Each `pNext` member of any structure (including this one) in the `pNext` chain **must** be either `NULL` or a pointer to a valid instance of [VkMemoryDedicatedRequirements](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryDedicatedRequirements.html) or [VkTileMemoryRequirementsQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTileMemoryRequirementsQCOM.html)
- [](#VUID-VkMemoryRequirements2-sType-unique)VUID-VkMemoryRequirements2-sType-unique  
  The `sType` value of each structure in the `pNext` chain **must** be unique

## [](#_see_also)See Also

[VK\_VERSION\_1\_1](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_1.html), [VkMemoryRequirements](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryRequirements.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkGetBufferMemoryRequirements2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetBufferMemoryRequirements2.html), [vkGetBufferMemoryRequirements2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetBufferMemoryRequirements2.html), [vkGetDataGraphPipelineSessionMemoryRequirementsARM](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetDataGraphPipelineSessionMemoryRequirementsARM.html), [vkGetDeviceBufferMemoryRequirements](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetDeviceBufferMemoryRequirements.html), [vkGetDeviceBufferMemoryRequirements](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetDeviceBufferMemoryRequirements.html), [vkGetDeviceImageMemoryRequirements](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetDeviceImageMemoryRequirements.html), [vkGetDeviceImageMemoryRequirements](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetDeviceImageMemoryRequirements.html), [vkGetDeviceTensorMemoryRequirementsARM](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetDeviceTensorMemoryRequirementsARM.html), [vkGetGeneratedCommandsMemoryRequirementsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetGeneratedCommandsMemoryRequirementsEXT.html), [vkGetGeneratedCommandsMemoryRequirementsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetGeneratedCommandsMemoryRequirementsNV.html), [vkGetImageMemoryRequirements2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetImageMemoryRequirements2.html), [vkGetImageMemoryRequirements2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetImageMemoryRequirements2.html), [vkGetPipelineIndirectMemoryRequirementsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPipelineIndirectMemoryRequirementsNV.html), [vkGetTensorMemoryRequirementsARM](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetTensorMemoryRequirementsARM.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkMemoryRequirements2).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0