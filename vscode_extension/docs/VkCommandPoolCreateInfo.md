# VkCommandPoolCreateInfo(3) Manual Page

## Name

VkCommandPoolCreateInfo - Structure specifying parameters of a newly created command pool



## [](#_c_specification)C Specification

The `VkCommandPoolCreateInfo` structure is defined as:

```c++
// Provided by VK_VERSION_1_0
typedef struct VkCommandPoolCreateInfo {
    VkStructureType             sType;
    const void*                 pNext;
    VkCommandPoolCreateFlags    flags;
    uint32_t                    queueFamilyIndex;
} VkCommandPoolCreateInfo;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `flags` is a bitmask of [VkCommandPoolCreateFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandPoolCreateFlagBits.html) indicating usage behavior for the pool and command buffers allocated from it.
- `queueFamilyIndex` designates a queue family as described in section [Queue Family Properties](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#devsandqueues-queueprops). All command buffers allocated from this command pool **must** be submitted on queues from the same queue family.

## [](#_description)Description

Valid Usage

- [](#VUID-VkCommandPoolCreateInfo-flags-02860)VUID-VkCommandPoolCreateInfo-flags-02860  
  If the [`protectedMemory`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-protectedMemory) feature is not enabled, the `VK_COMMAND_POOL_CREATE_PROTECTED_BIT` bit of `flags` **must** not be set
- [](#VUID-VkCommandPoolCreateInfo-pNext-09908)VUID-VkCommandPoolCreateInfo-pNext-09908  
  If the `pNext` chain includes a [VkDataGraphProcessingEngineCreateInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphProcessingEngineCreateInfoARM.html) structure, then `queueFamilyIndex` **must** designate a queue family that supports `VK_QUEUE_DATA_GRAPH_BIT_ARM`
- [](#VUID-VkCommandPoolCreateInfo-pNext-09909)VUID-VkCommandPoolCreateInfo-pNext-09909  
  If the `pNext` chain includes a [VkDataGraphProcessingEngineCreateInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphProcessingEngineCreateInfoARM.html) structure, each member of `pProcessingEngines` **must** be identical to [VkQueueFamilyDataGraphPropertiesARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueueFamilyDataGraphPropertiesARM.html)::`engine` retrieved from [vkGetPhysicalDeviceQueueFamilyDataGraphPropertiesARM](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceQueueFamilyDataGraphPropertiesARM.html) with `queueFamilyIndex` and the `physicalDevice` that was used to create `device`

Valid Usage (Implicit)

- [](#VUID-VkCommandPoolCreateInfo-sType-sType)VUID-VkCommandPoolCreateInfo-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_COMMAND_POOL_CREATE_INFO`
- [](#VUID-VkCommandPoolCreateInfo-pNext-pNext)VUID-VkCommandPoolCreateInfo-pNext-pNext  
  `pNext` **must** be `NULL` or a pointer to a valid instance of [VkDataGraphProcessingEngineCreateInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphProcessingEngineCreateInfoARM.html)
- [](#VUID-VkCommandPoolCreateInfo-sType-unique)VUID-VkCommandPoolCreateInfo-sType-unique  
  The `sType` value of each structure in the `pNext` chain **must** be unique
- [](#VUID-VkCommandPoolCreateInfo-flags-parameter)VUID-VkCommandPoolCreateInfo-flags-parameter  
  `flags` **must** be a valid combination of [VkCommandPoolCreateFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandPoolCreateFlagBits.html) values

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkCommandPoolCreateFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandPoolCreateFlags.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkCreateCommandPool](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateCommandPool.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkCommandPoolCreateInfo).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0