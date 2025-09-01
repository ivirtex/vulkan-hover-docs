# VkQueueFamilyDataGraphProcessingEnginePropertiesARM(3) Manual Page

## Name

VkQueueFamilyDataGraphProcessingEnginePropertiesARM - Structure describing the properties of a data graph processing engine type for a given queue family



## [](#_c_specification)C Specification

The `VkQueueFamilyDataGraphProcessingEnginePropertiesARM` structure is defined as:

```c++
// Provided by VK_ARM_data_graph
typedef struct VkQueueFamilyDataGraphProcessingEnginePropertiesARM {
    VkStructureType                       sType;
    const void*                           pNext;
    VkExternalSemaphoreHandleTypeFlags    foreignSemaphoreHandleTypes;
    VkExternalMemoryHandleTypeFlags       foreignMemoryHandleTypes;
} VkQueueFamilyDataGraphProcessingEnginePropertiesARM;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `foreignSemaphoreHandleTypes` is a [VkExternalSemaphoreHandleTypeFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalSemaphoreHandleTypeFlags.html) that describes the external semaphore handle types supported by a foreign data graph processing engine.
- `foreignMemoryHandleTypes` is a [VkExternalMemoryHandleTypeFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalMemoryHandleTypeFlags.html) that describes the external memory handle types supported by a foreign data graph processing engine.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkQueueFamilyDataGraphProcessingEnginePropertiesARM-sType-sType)VUID-VkQueueFamilyDataGraphProcessingEnginePropertiesARM-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_QUEUE_FAMILY_DATA_GRAPH_PROCESSING_ENGINE_PROPERTIES_ARM`
- [](#VUID-VkQueueFamilyDataGraphProcessingEnginePropertiesARM-pNext-pNext)VUID-VkQueueFamilyDataGraphProcessingEnginePropertiesARM-pNext-pNext  
  `pNext` **must** be `NULL`
- [](#VUID-VkQueueFamilyDataGraphProcessingEnginePropertiesARM-foreignSemaphoreHandleTypes-parameter)VUID-VkQueueFamilyDataGraphProcessingEnginePropertiesARM-foreignSemaphoreHandleTypes-parameter  
  `foreignSemaphoreHandleTypes` **must** be a valid combination of [VkExternalSemaphoreHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalSemaphoreHandleTypeFlagBits.html) values
- [](#VUID-VkQueueFamilyDataGraphProcessingEnginePropertiesARM-foreignSemaphoreHandleTypes-requiredbitmask)VUID-VkQueueFamilyDataGraphProcessingEnginePropertiesARM-foreignSemaphoreHandleTypes-requiredbitmask  
  `foreignSemaphoreHandleTypes` **must** not be `0`
- [](#VUID-VkQueueFamilyDataGraphProcessingEnginePropertiesARM-foreignMemoryHandleTypes-parameter)VUID-VkQueueFamilyDataGraphProcessingEnginePropertiesARM-foreignMemoryHandleTypes-parameter  
  `foreignMemoryHandleTypes` **must** be a valid combination of [VkExternalMemoryHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalMemoryHandleTypeFlagBits.html) values
- [](#VUID-VkQueueFamilyDataGraphProcessingEnginePropertiesARM-foreignMemoryHandleTypes-requiredbitmask)VUID-VkQueueFamilyDataGraphProcessingEnginePropertiesARM-foreignMemoryHandleTypes-requiredbitmask  
  `foreignMemoryHandleTypes` **must** not be `0`

## [](#_see_also)See Also

[VK\_ARM\_data\_graph](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_ARM_data_graph.html), [VkExternalMemoryHandleTypeFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalMemoryHandleTypeFlags.html), [VkExternalSemaphoreHandleTypeFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalSemaphoreHandleTypeFlags.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkGetPhysicalDeviceQueueFamilyDataGraphProcessingEnginePropertiesARM](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceQueueFamilyDataGraphProcessingEnginePropertiesARM.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkQueueFamilyDataGraphProcessingEnginePropertiesARM).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0