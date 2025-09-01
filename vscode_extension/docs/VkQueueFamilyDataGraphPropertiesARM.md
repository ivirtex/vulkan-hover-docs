# VkQueueFamilyDataGraphPropertiesARM(3) Manual Page

## Name

VkQueueFamilyDataGraphPropertiesARM - Structure describing a data graph processing engine and operation it supports



## [](#_c_specification)C Specification

The `VkQueueFamilyDataGraphPropertiesARM` structure is defined as:

```c++
// Provided by VK_ARM_data_graph
typedef struct VkQueueFamilyDataGraphPropertiesARM {
    VkStructureType                                 sType;
    const void*                                     pNext;
    VkPhysicalDeviceDataGraphProcessingEngineARM    engine;
    VkPhysicalDeviceDataGraphOperationSupportARM    operation;
} VkQueueFamilyDataGraphPropertiesARM;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `engine` is a [VkPhysicalDeviceDataGraphProcessingEngineARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceDataGraphProcessingEngineARM.html) structure describing a data graph processing engine.
- `operation` is a [VkPhysicalDeviceDataGraphOperationSupportARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceDataGraphOperationSupportARM.html) structure describing one or more operations supported by a data graph processing engine.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkQueueFamilyDataGraphPropertiesARM-sType-sType)VUID-VkQueueFamilyDataGraphPropertiesARM-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_QUEUE_FAMILY_DATA_GRAPH_PROPERTIES_ARM`
- [](#VUID-VkQueueFamilyDataGraphPropertiesARM-pNext-pNext)VUID-VkQueueFamilyDataGraphPropertiesARM-pNext-pNext  
  `pNext` **must** be `NULL`
- [](#VUID-VkQueueFamilyDataGraphPropertiesARM-engine-parameter)VUID-VkQueueFamilyDataGraphPropertiesARM-engine-parameter  
  `engine` **must** be a valid [VkPhysicalDeviceDataGraphProcessingEngineARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceDataGraphProcessingEngineARM.html) structure
- [](#VUID-VkQueueFamilyDataGraphPropertiesARM-operation-parameter)VUID-VkQueueFamilyDataGraphPropertiesARM-operation-parameter  
  `operation` **must** be a valid [VkPhysicalDeviceDataGraphOperationSupportARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceDataGraphOperationSupportARM.html) structure

## [](#_see_also)See Also

[VK\_ARM\_data\_graph](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_ARM_data_graph.html), [VkPhysicalDeviceDataGraphOperationSupportARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceDataGraphOperationSupportARM.html), [VkPhysicalDeviceDataGraphProcessingEngineARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceDataGraphProcessingEngineARM.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkGetPhysicalDeviceQueueFamilyDataGraphPropertiesARM](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceQueueFamilyDataGraphPropertiesARM.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkQueueFamilyDataGraphPropertiesARM).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0