# VkPipelineIndirectDeviceAddressInfoNV(3) Manual Page

## Name

VkPipelineIndirectDeviceAddressInfoNV - Structure specifying the pipeline to query an address for



## [](#_c_specification)C Specification

The `VkPipelineIndirectDeviceAddressInfoNV` structure is defined as:

```c++
// Provided by VK_NV_device_generated_commands_compute
typedef struct VkPipelineIndirectDeviceAddressInfoNV {
    VkStructureType        sType;
    const void*            pNext;
    VkPipelineBindPoint    pipelineBindPoint;
    VkPipeline             pipeline;
} VkPipelineIndirectDeviceAddressInfoNV;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `pipelineBindPoint` is a [VkPipelineBindPoint](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineBindPoint.html) value specifying the type of pipeline whose device address is being queried.
- `pipeline` specifies the pipeline whose device address is being queried.

## [](#_description)Description

Valid Usage

- [](#VUID-VkPipelineIndirectDeviceAddressInfoNV-pipelineBindPoint-09079)VUID-VkPipelineIndirectDeviceAddressInfoNV-pipelineBindPoint-09079  
  The provided `pipelineBindPoint` **must** be of type `VK_PIPELINE_BIND_POINT_COMPUTE`
- [](#VUID-VkPipelineIndirectDeviceAddressInfoNV-pipeline-09080)VUID-VkPipelineIndirectDeviceAddressInfoNV-pipeline-09080  
  `pipeline` **must** have been created with flag `VK_PIPELINE_CREATE_INDIRECT_BINDABLE_BIT_NV` set
- [](#VUID-VkPipelineIndirectDeviceAddressInfoNV-pipeline-09081)VUID-VkPipelineIndirectDeviceAddressInfoNV-pipeline-09081  
  `pipeline` **must** have been created with a [VkComputePipelineIndirectBufferInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkComputePipelineIndirectBufferInfoNV.html) structure specifying a valid address where its metadata will be saved

Valid Usage (Implicit)

- [](#VUID-VkPipelineIndirectDeviceAddressInfoNV-sType-sType)VUID-VkPipelineIndirectDeviceAddressInfoNV-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PIPELINE_INDIRECT_DEVICE_ADDRESS_INFO_NV`
- [](#VUID-VkPipelineIndirectDeviceAddressInfoNV-pNext-pNext)VUID-VkPipelineIndirectDeviceAddressInfoNV-pNext-pNext  
  `pNext` **must** be `NULL`
- [](#VUID-VkPipelineIndirectDeviceAddressInfoNV-pipelineBindPoint-parameter)VUID-VkPipelineIndirectDeviceAddressInfoNV-pipelineBindPoint-parameter  
  `pipelineBindPoint` **must** be a valid [VkPipelineBindPoint](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineBindPoint.html) value
- [](#VUID-VkPipelineIndirectDeviceAddressInfoNV-pipeline-parameter)VUID-VkPipelineIndirectDeviceAddressInfoNV-pipeline-parameter  
  `pipeline` **must** be a valid [VkPipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipeline.html) handle

## [](#_see_also)See Also

[VK\_NV\_device\_generated\_commands\_compute](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_device_generated_commands_compute.html), [VkPipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipeline.html), [VkPipelineBindPoint](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineBindPoint.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkGetPipelineIndirectDeviceAddressNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPipelineIndirectDeviceAddressNV.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPipelineIndirectDeviceAddressInfoNV)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0