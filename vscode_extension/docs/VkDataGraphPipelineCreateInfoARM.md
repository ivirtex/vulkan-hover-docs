# VkDataGraphPipelineCreateInfoARM(3) Manual Page

## Name

VkDataGraphPipelineCreateInfoARM - Structure specifying parameters of a newly created data graph pipeline



## [](#_c_specification)C Specification

The `VkDataGraphPipelineCreateInfoARM` structure is defined as:

```c++
// Provided by VK_ARM_data_graph
typedef struct VkDataGraphPipelineCreateInfoARM {
    VkStructureType                              sType;
    const void*                                  pNext;
    VkPipelineCreateFlags2KHR                    flags;
    VkPipelineLayout                             layout;
    uint32_t                                     resourceInfoCount;
    const VkDataGraphPipelineResourceInfoARM*    pResourceInfos;
} VkDataGraphPipelineCreateInfoARM;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `flags` is a bitmask of [VkPipelineCreateFlagBits2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCreateFlagBits2KHR.html) specifying how the pipeline will be generated.
- `layout` is the description of binding locations used by both the pipeline and descriptor sets used with the pipeline.
- `resourceInfoCount` is the length of the `pResourceInfos` array.
- `pResourceInfos` is a pointer to an array of [VkDataGraphPipelineResourceInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphPipelineResourceInfoARM.html) structures.

## [](#_description)Description

Applications **can** create a data graph pipeline entirely from data present in a pipeline cache. This is done by including a [VkDataGraphPipelineIdentifierCreateInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphPipelineIdentifierCreateInfoARM.html) structure in the `pNext` chain. If the required data is not found in the pipeline cache, creating the data graph pipeline is not possible and the implementation **must** fail as specified by `VK_PIPELINE_CREATE_FAIL_ON_PIPELINE_COMPILE_REQUIRED_BIT`.

When an identifier is used to create a data graph pipeline, implementations **may** fail pipeline creation with `VK_PIPELINE_COMPILE_REQUIRED` for any reason.

The data graph engines for this pipeline **can** be selected by including a [VkDataGraphProcessingEngineCreateInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphProcessingEngineCreateInfoARM.html) to the `pNext` chain of this structure. Otherwise, `VK_PHYSICAL_DEVICE_DATA_GRAPH_PROCESSING_ENGINE_TYPE_DEFAULT_ARM` will be used as the sole data graph engine.

The data graph operations that this pipeline uses **must** be supported for the data graph engines selected for this pipeline as retrieved by [vkGetPhysicalDeviceQueueFamilyDataGraphPropertiesARM](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceQueueFamilyDataGraphPropertiesARM.html).

Valid Usage

- [](#VUID-VkDataGraphPipelineCreateInfoARM-pNext-09763)VUID-VkDataGraphPipelineCreateInfoARM-pNext-09763  
  One and only one of the following structures **must** be included in the `pNext` chain:
  
  - [VkDataGraphPipelineShaderModuleCreateInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphPipelineShaderModuleCreateInfoARM.html)
  - [VkDataGraphPipelineIdentifierCreateInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphPipelineIdentifierCreateInfoARM.html)
- [](#VUID-VkDataGraphPipelineCreateInfoARM-flags-09764)VUID-VkDataGraphPipelineCreateInfoARM-flags-09764  
  `flags` **may** only contain `VK_PIPELINE_CREATE_2_NO_PROTECTED_ACCESS_BIT_EXT`, `VK_PIPELINE_CREATE_2_PROTECTED_ACCESS_ONLY_BIT_EXT`, `VK_PIPELINE_CREATE_2_DISABLE_OPTIMIZATION_BIT`, `VK_PIPELINE_CREATE_2_DESCRIPTOR_BUFFER_BIT_EXT`, `VK_PIPELINE_CREATE_2_FAIL_ON_PIPELINE_COMPILE_REQUIRED_BIT_KHR` or `VK_PIPELINE_CREATE_2_EARLY_RETURN_ON_FAILURE_BIT_KHR`
- [](#VUID-VkDataGraphPipelineCreateInfoARM-layout-09767)VUID-VkDataGraphPipelineCreateInfoARM-layout-09767  
  `layout` **must** have been created with `pushConstantRangeCount` equal to 0 and `pPushConstantRanges` equal to `NULL`
- [](#VUID-VkDataGraphPipelineCreateInfoARM-dataGraphUpdateAfterBind-09768)VUID-VkDataGraphPipelineCreateInfoARM-dataGraphUpdateAfterBind-09768  
  If the [`dataGraphUpdateAfterBind`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-dataGraphUpdateAfterBind) feature is not enabled, `layout` must not use any [VkDescriptorSetLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorSetLayout.html) object created with the `VK_DESCRIPTOR_SET_LAYOUT_CREATE_UPDATE_AFTER_BIND_POOL_BIT` bit set
- [](#VUID-VkDataGraphPipelineCreateInfoARM-dataGraphDescriptorBuffer-09885)VUID-VkDataGraphPipelineCreateInfoARM-dataGraphDescriptorBuffer-09885  
  If the [`dataGraphDescriptorBuffer`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-dataGraphDescriptorBuffer) feature is not enabled, `flags` **must** not contain `VK_PIPELINE_CREATE_2_DESCRIPTOR_BUFFER_BIT_EXT`
- [](#VUID-VkDataGraphPipelineCreateInfoARM-layout-09769)VUID-VkDataGraphPipelineCreateInfoARM-layout-09769  
  If a [VkDataGraphPipelineShaderModuleCreateInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphPipelineShaderModuleCreateInfoARM.html) structure is included in the `pNext` chain and a [resource variable](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#interfaces-resources) is declared in the shader module, the corresponding descriptor binding used to create `layout` **must** have a `descriptorType` that corresponds to the type of the [resource variable](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#interfaces-resources)
- [](#VUID-VkDataGraphPipelineCreateInfoARM-pNext-09875)VUID-VkDataGraphPipelineCreateInfoARM-pNext-09875  
  If a [VkDataGraphPipelineIdentifierCreateInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphPipelineIdentifierCreateInfoARM.html) structure is included in the `pNext` chain, then `flags` **must** contain `VK_PIPELINE_CREATE_2_FAIL_ON_PIPELINE_COMPILE_REQUIRED_BIT`
- [](#VUID-VkDataGraphPipelineCreateInfoARM-pNext-09882)VUID-VkDataGraphPipelineCreateInfoARM-pNext-09882  
  If a [VkDataGraphPipelineIdentifierCreateInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphPipelineIdentifierCreateInfoARM.html) structure is included in the `pNext` chain, then `resourceInfoCount` **must** be 0 and `pResourceInfos` **must** be `NULL`
- [](#VUID-VkDataGraphPipelineCreateInfoARM-dataGraphShaderModule-09886)VUID-VkDataGraphPipelineCreateInfoARM-dataGraphShaderModule-09886  
  If the [`dataGraphShaderModule`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-dataGraphShaderModule) feature is not enabled, a [VkDataGraphPipelineShaderModuleCreateInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphPipelineShaderModuleCreateInfoARM.html) structure **must** not be included in the `pNext` chain
- [](#VUID-VkDataGraphPipelineCreateInfoARM-layout-09934)VUID-VkDataGraphPipelineCreateInfoARM-layout-09934  
  If a [VkDataGraphPipelineShaderModuleCreateInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphPipelineShaderModuleCreateInfoARM.html) structure is included in the `pNext` chain and an array [resource variable](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#interfaces-resources) is declared in the shader module, the corresponding descriptor binding used to create `layout` **must** have a `descriptorCount` that matches the length of the array
- [](#VUID-VkDataGraphPipelineCreateInfoARM-pipelineCreationCacheControl-09871)VUID-VkDataGraphPipelineCreateInfoARM-pipelineCreationCacheControl-09871  
  If the [`pipelineCreationCacheControl`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-pipelineCreationCacheControl) feature is not enabled, `flags` **must** not include `VK_PIPELINE_CREATE_2_FAIL_ON_PIPELINE_COMPILE_REQUIRED_BIT_KHR` or `VK_PIPELINE_CREATE_2_EARLY_RETURN_ON_FAILURE_BIT_KHR`
- [](#VUID-VkDataGraphPipelineCreateInfoARM-pSetLayouts-09770)VUID-VkDataGraphPipelineCreateInfoARM-pSetLayouts-09770  
  The descriptor set layouts in [VkPipelineLayoutCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineLayoutCreateInfo.html)::`pSetLayouts` used to create `layout` **must** not include any [VkDescriptorSetLayoutBinding](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorSetLayoutBinding.html) whose descriptor type is `VK_DESCRIPTOR_TYPE_MUTABLE_EXT`
- [](#VUID-VkDataGraphPipelineCreateInfoARM-pipelineProtectedAccess-09772)VUID-VkDataGraphPipelineCreateInfoARM-pipelineProtectedAccess-09772  
  If the [`pipelineProtectedAccess`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-pipelineProtectedAccess) feature is not enabled, `flags` **must** not include `VK_PIPELINE_CREATE_2_NO_PROTECTED_ACCESS_BIT_EXT` or `VK_PIPELINE_CREATE_2_PROTECTED_ACCESS_ONLY_BIT_EXT`
- [](#VUID-VkDataGraphPipelineCreateInfoARM-flags-09773)VUID-VkDataGraphPipelineCreateInfoARM-flags-09773  
  `flags` **must** not include both `VK_PIPELINE_CREATE_2_NO_PROTECTED_ACCESS_BIT_EXT` and `VK_PIPELINE_CREATE_2_PROTECTED_ACCESS_ONLY_BIT_EXT`
- [](#VUID-VkDataGraphPipelineCreateInfoARM-pNext-09804)VUID-VkDataGraphPipelineCreateInfoARM-pNext-09804  
  If the `pNext` chain includes an [VkPipelineCreationFeedbackCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCreationFeedbackCreateInfo.html) structure, then its `pipelineStageCreationFeedbackCount` **must** be 0
- [](#VUID-VkDataGraphPipelineCreateInfoARM-pNext-09948)VUID-VkDataGraphPipelineCreateInfoARM-pNext-09948  
  If a [VkDataGraphProcessingEngineCreateInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphProcessingEngineCreateInfoARM.html) structure is included in the `pNext` chain, each member of `pProcessingEngines` **must** be identical to an [VkQueueFamilyDataGraphPropertiesARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueueFamilyDataGraphPropertiesARM.html)::`engine` retrieved from [vkGetPhysicalDeviceQueueFamilyDataGraphPropertiesARM](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceQueueFamilyDataGraphPropertiesARM.html) with the `physicalDevice` that was used to create `device`
- [](#VUID-VkDataGraphPipelineCreateInfoARM-pNext-09949)VUID-VkDataGraphPipelineCreateInfoARM-pNext-09949  
  If a [VkDataGraphProcessingEngineCreateInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphProcessingEngineCreateInfoARM.html) structure is not included in the `pNext` chain, `VK_PHYSICAL_DEVICE_DATA_GRAPH_PROCESSING_ENGINE_TYPE_DEFAULT_ARM` **must** be set in an [VkQueueFamilyDataGraphPropertiesARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueueFamilyDataGraphPropertiesARM.html)::`engine` retrieved from [vkGetPhysicalDeviceQueueFamilyDataGraphPropertiesARM](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceQueueFamilyDataGraphPropertiesARM.html) with the `physicalDevice` that was used to create `device`

Valid Usage (Implicit)

- [](#VUID-VkDataGraphPipelineCreateInfoARM-sType-sType)VUID-VkDataGraphPipelineCreateInfoARM-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_DATA_GRAPH_PIPELINE_CREATE_INFO_ARM`
- [](#VUID-VkDataGraphPipelineCreateInfoARM-pNext-pNext)VUID-VkDataGraphPipelineCreateInfoARM-pNext-pNext  
  Each `pNext` member of any structure (including this one) in the `pNext` chain **must** be either `NULL` or a pointer to a valid instance of [VkDataGraphPipelineCompilerControlCreateInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphPipelineCompilerControlCreateInfoARM.html), [VkDataGraphPipelineIdentifierCreateInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphPipelineIdentifierCreateInfoARM.html), [VkDataGraphPipelineShaderModuleCreateInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphPipelineShaderModuleCreateInfoARM.html), [VkDataGraphProcessingEngineCreateInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphProcessingEngineCreateInfoARM.html), [VkPipelineCreationFeedbackCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCreationFeedbackCreateInfo.html), or [VkShaderModuleCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderModuleCreateInfo.html)
- [](#VUID-VkDataGraphPipelineCreateInfoARM-sType-unique)VUID-VkDataGraphPipelineCreateInfoARM-sType-unique  
  The `sType` value of each structure in the `pNext` chain **must** be unique
- [](#VUID-VkDataGraphPipelineCreateInfoARM-flags-parameter)VUID-VkDataGraphPipelineCreateInfoARM-flags-parameter  
  `flags` **must** be a valid combination of [VkPipelineCreateFlagBits2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCreateFlagBits2KHR.html) values
- [](#VUID-VkDataGraphPipelineCreateInfoARM-layout-parameter)VUID-VkDataGraphPipelineCreateInfoARM-layout-parameter  
  `layout` **must** be a valid [VkPipelineLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineLayout.html) handle
- [](#VUID-VkDataGraphPipelineCreateInfoARM-pResourceInfos-parameter)VUID-VkDataGraphPipelineCreateInfoARM-pResourceInfos-parameter  
  `pResourceInfos` **must** be a valid pointer to an array of `resourceInfoCount` valid [VkDataGraphPipelineResourceInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphPipelineResourceInfoARM.html) structures
- [](#VUID-VkDataGraphPipelineCreateInfoARM-resourceInfoCount-arraylength)VUID-VkDataGraphPipelineCreateInfoARM-resourceInfoCount-arraylength  
  `resourceInfoCount` **must** be greater than `0`

## [](#_see_also)See Also

[VK\_ARM\_data\_graph](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_ARM_data_graph.html), [VkDataGraphPipelineResourceInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphPipelineResourceInfoARM.html), [VkPipelineCreateFlags2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCreateFlags2.html), [VkPipelineLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineLayout.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkCreateDataGraphPipelinesARM](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateDataGraphPipelinesARM.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkDataGraphPipelineCreateInfoARM).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0