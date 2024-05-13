# VkPipelineIndirectDeviceAddressInfoNV(3) Manual Page

## Name

VkPipelineIndirectDeviceAddressInfoNV - Structure specifying the
pipeline to query an address for



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPipelineIndirectDeviceAddressInfoNV` structure is defined as:

``` c
// Provided by VK_NV_device_generated_commands_compute
typedef struct VkPipelineIndirectDeviceAddressInfoNV {
    VkStructureType        sType;
    const void*            pNext;
    VkPipelineBindPoint    pipelineBindPoint;
    VkPipeline             pipeline;
} VkPipelineIndirectDeviceAddressInfoNV;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `pipelineBindPoint` is a
  [VkPipelineBindPoint](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineBindPoint.html) value specifying the
  type of pipeline whose device address is being queried.

- `pipeline` specifies the pipeline whose device address is being
  queried.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a
  href="#VUID-VkPipelineIndirectDeviceAddressInfoNV-pipelineBindPoint-09079"
  id="VUID-VkPipelineIndirectDeviceAddressInfoNV-pipelineBindPoint-09079"></a>
  VUID-VkPipelineIndirectDeviceAddressInfoNV-pipelineBindPoint-09079  
  The provided `pipelineBindPoint` **must** be of type
  `VK_PIPELINE_BIND_POINT_COMPUTE`

- <a href="#VUID-VkPipelineIndirectDeviceAddressInfoNV-pipeline-09080"
  id="VUID-VkPipelineIndirectDeviceAddressInfoNV-pipeline-09080"></a>
  VUID-VkPipelineIndirectDeviceAddressInfoNV-pipeline-09080  
  `pipeline` **must** have been created with flag
  `VK_PIPELINE_CREATE_INDIRECT_BINDABLE_BIT_NV` set

- <a href="#VUID-VkPipelineIndirectDeviceAddressInfoNV-pipeline-09081"
  id="VUID-VkPipelineIndirectDeviceAddressInfoNV-pipeline-09081"></a>
  VUID-VkPipelineIndirectDeviceAddressInfoNV-pipeline-09081  
  `pipeline` **must** have been created with a
  [VkComputePipelineIndirectBufferInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkComputePipelineIndirectBufferInfoNV.html)
  structure specifying a valid address where its metadata will be saved

Valid Usage (Implicit)

- <a href="#VUID-VkPipelineIndirectDeviceAddressInfoNV-sType-sType"
  id="VUID-VkPipelineIndirectDeviceAddressInfoNV-sType-sType"></a>
  VUID-VkPipelineIndirectDeviceAddressInfoNV-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PIPELINE_INDIRECT_DEVICE_ADDRESS_INFO_NV`

- <a href="#VUID-VkPipelineIndirectDeviceAddressInfoNV-pNext-pNext"
  id="VUID-VkPipelineIndirectDeviceAddressInfoNV-pNext-pNext"></a>
  VUID-VkPipelineIndirectDeviceAddressInfoNV-pNext-pNext  
  `pNext` **must** be `NULL`

- <a
  href="#VUID-VkPipelineIndirectDeviceAddressInfoNV-pipelineBindPoint-parameter"
  id="VUID-VkPipelineIndirectDeviceAddressInfoNV-pipelineBindPoint-parameter"></a>
  VUID-VkPipelineIndirectDeviceAddressInfoNV-pipelineBindPoint-parameter  
  `pipelineBindPoint` **must** be a valid
  [VkPipelineBindPoint](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineBindPoint.html) value

- <a href="#VUID-VkPipelineIndirectDeviceAddressInfoNV-pipeline-parameter"
  id="VUID-VkPipelineIndirectDeviceAddressInfoNV-pipeline-parameter"></a>
  VUID-VkPipelineIndirectDeviceAddressInfoNV-pipeline-parameter  
  `pipeline` **must** be a valid [VkPipeline](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipeline.html) handle

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_device_generated_commands_compute](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_device_generated_commands_compute.html),
[VkPipeline](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipeline.html),
[VkPipelineBindPoint](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineBindPoint.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkGetPipelineIndirectDeviceAddressNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPipelineIndirectDeviceAddressNV.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPipelineIndirectDeviceAddressInfoNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
