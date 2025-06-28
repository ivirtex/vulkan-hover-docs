# VkDataGraphProcessingEngineCreateInfoARM(3) Manual Page

## Name

VkDataGraphProcessingEngineCreateInfoARM - Structure describing a collection of data graph processing engines for which the object being created is specialized



## [](#_c_specification)C Specification

The `VkDataGraphProcessingEngineCreateInfoARM` structure is defined as:

```c++
// Provided by VK_ARM_data_graph
typedef struct VkDataGraphProcessingEngineCreateInfoARM {
    VkStructureType                                  sType;
    const void*                                      pNext;
    uint32_t                                         processingEngineCount;
    VkPhysicalDeviceDataGraphProcessingEngineARM*    pProcessingEngines;
} VkDataGraphProcessingEngineCreateInfoARM;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `processingEngineCount` is the number of elements in `pProcessingEngines`.
- `pProcessingEngines` is a pointer to an array of `processingEngineCount` [VkPhysicalDeviceDataGraphProcessingEngineARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceDataGraphProcessingEngineARM.html) structures.

## [](#_description)Description

Valid Usage

- [](#VUID-VkDataGraphProcessingEngineCreateInfoARM-pProcessingEngines-09918)VUID-VkDataGraphProcessingEngineCreateInfoARM-pProcessingEngines-09918  
  `pProcessingEngines` **must** not contain identical [VkPhysicalDeviceDataGraphProcessingEngineARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceDataGraphProcessingEngineARM.html) structures

Valid Usage (Implicit)

- [](#VUID-VkDataGraphProcessingEngineCreateInfoARM-sType-sType)VUID-VkDataGraphProcessingEngineCreateInfoARM-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_DATA_GRAPH_PROCESSING_ENGINE_CREATE_INFO_ARM`
- [](#VUID-VkDataGraphProcessingEngineCreateInfoARM-pProcessingEngines-parameter)VUID-VkDataGraphProcessingEngineCreateInfoARM-pProcessingEngines-parameter  
  `pProcessingEngines` **must** be a valid pointer to an array of `processingEngineCount` [VkPhysicalDeviceDataGraphProcessingEngineARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceDataGraphProcessingEngineARM.html) structures
- [](#VUID-VkDataGraphProcessingEngineCreateInfoARM-processingEngineCount-arraylength)VUID-VkDataGraphProcessingEngineCreateInfoARM-processingEngineCount-arraylength  
  `processingEngineCount` **must** be greater than `0`

## [](#_see_also)See Also

[VK\_ARM\_data\_graph](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_ARM_data_graph.html), [VkPhysicalDeviceDataGraphProcessingEngineARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceDataGraphProcessingEngineARM.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkDataGraphProcessingEngineCreateInfoARM)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0