# VkReleaseCapturedPipelineDataInfoKHR(3) Manual Page

## Name

VkReleaseCapturedPipelineDataInfoKHR - Structure specifying a pipeline whose captured data is to be released



## [](#_c_specification)C Specification

The `VkReleaseCapturedPipelineDataInfoKHR` structure is defined as:

```c++
// Provided by VK_KHR_pipeline_binary
typedef struct VkReleaseCapturedPipelineDataInfoKHR {
    VkStructureType    sType;
    void*              pNext;
    VkPipeline         pipeline;
} VkReleaseCapturedPipelineDataInfoKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `pipeline` the handle of the pipeline object to release the data from.

## [](#_description)Description

Valid Usage

- [](#VUID-VkReleaseCapturedPipelineDataInfoKHR-pipeline-09613)VUID-VkReleaseCapturedPipelineDataInfoKHR-pipeline-09613  
  `pipeline` **must** have been created with `VK_PIPELINE_CREATE_2_CAPTURE_DATA_BIT_KHR`
- [](#VUID-VkReleaseCapturedPipelineDataInfoKHR-pipeline-09618)VUID-VkReleaseCapturedPipelineDataInfoKHR-pipeline-09618  
  `pipeline` **must** not have been used in a previous call to `vkReleaseCapturedPipelineDataKHR`

Valid Usage (Implicit)

- [](#VUID-VkReleaseCapturedPipelineDataInfoKHR-sType-sType)VUID-VkReleaseCapturedPipelineDataInfoKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_RELEASE_CAPTURED_PIPELINE_DATA_INFO_KHR`
- [](#VUID-VkReleaseCapturedPipelineDataInfoKHR-pNext-pNext)VUID-VkReleaseCapturedPipelineDataInfoKHR-pNext-pNext  
  `pNext` **must** be `NULL`
- [](#VUID-VkReleaseCapturedPipelineDataInfoKHR-pipeline-parameter)VUID-VkReleaseCapturedPipelineDataInfoKHR-pipeline-parameter  
  `pipeline` **must** be a valid [VkPipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipeline.html) handle

Host Synchronization

- Host access to `pipeline` **must** be externally synchronized

## [](#_see_also)See Also

[VK\_KHR\_pipeline\_binary](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_pipeline_binary.html), [VkPipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipeline.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkReleaseCapturedPipelineDataKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkReleaseCapturedPipelineDataKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkReleaseCapturedPipelineDataInfoKHR).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0