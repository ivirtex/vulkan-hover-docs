# VkPhysicalDeviceQueueFamilyDataGraphProcessingEngineInfoARM(3) Manual Page

## Name

VkPhysicalDeviceQueueFamilyDataGraphProcessingEngineInfoARM - Structure specifying a data graph processing engine type and queue family to query



## [](#_c_specification)C Specification

The `VkPhysicalDeviceQueueFamilyDataGraphProcessingEngineInfoARM` structure is defined as:

```c++
// Provided by VK_ARM_data_graph
typedef struct VkPhysicalDeviceQueueFamilyDataGraphProcessingEngineInfoARM {
    VkStructureType                                     sType;
    const void*                                         pNext;
    uint32_t                                            queueFamilyIndex;
    VkPhysicalDeviceDataGraphProcessingEngineTypeARM    engineType;
} VkPhysicalDeviceQueueFamilyDataGraphProcessingEngineInfoARM;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `queueFamilyIndex` specifies the queue family being queried.
- `engineType` is a [VkPhysicalDeviceDataGraphProcessingEngineTypeARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceDataGraphProcessingEngineTypeARM.html) specifying the engine type whose properties are being queried.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDeviceQueueFamilyDataGraphProcessingEngineInfoARM-sType-sType)VUID-VkPhysicalDeviceQueueFamilyDataGraphProcessingEngineInfoARM-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_QUEUE_FAMILY_DATA_GRAPH_PROCESSING_ENGINE_INFO_ARM`
- [](#VUID-VkPhysicalDeviceQueueFamilyDataGraphProcessingEngineInfoARM-pNext-pNext)VUID-VkPhysicalDeviceQueueFamilyDataGraphProcessingEngineInfoARM-pNext-pNext  
  `pNext` **must** be `NULL`
- [](#VUID-VkPhysicalDeviceQueueFamilyDataGraphProcessingEngineInfoARM-engineType-parameter)VUID-VkPhysicalDeviceQueueFamilyDataGraphProcessingEngineInfoARM-engineType-parameter  
  `engineType` **must** be a valid [VkPhysicalDeviceDataGraphProcessingEngineTypeARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceDataGraphProcessingEngineTypeARM.html) value

## [](#_see_also)See Also

[VK\_ARM\_data\_graph](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_ARM_data_graph.html), [VkPhysicalDeviceDataGraphProcessingEngineTypeARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceDataGraphProcessingEngineTypeARM.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkGetPhysicalDeviceQueueFamilyDataGraphProcessingEnginePropertiesARM](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceQueueFamilyDataGraphProcessingEnginePropertiesARM.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceQueueFamilyDataGraphProcessingEngineInfoARM).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0