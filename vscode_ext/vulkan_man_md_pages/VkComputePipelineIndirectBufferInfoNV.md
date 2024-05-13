# VkComputePipelineIndirectBufferInfoNV(3) Manual Page

## Name

VkComputePipelineIndirectBufferInfoNV - Structure describing the device
address where pipeline's metadata will be saved



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkComputePipelineIndirectBufferInfoNV` structure is defined as:

``` c
// Provided by VK_NV_device_generated_commands_compute
typedef struct VkComputePipelineIndirectBufferInfoNV {
    VkStructureType    sType;
    const void*        pNext;
    VkDeviceAddress    deviceAddress;
    VkDeviceSize       size;
    VkDeviceAddress    pipelineDeviceAddressCaptureReplay;
} VkComputePipelineIndirectBufferInfoNV;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `deviceAddress` is the address where the pipeline’s metadata will be
  stored.

- `size` is the size of pipeline’s metadata that was queried using
  [vkGetPipelineIndirectMemoryRequirementsNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPipelineIndirectMemoryRequirementsNV.html).

- `pipelineDeviceAddressCaptureReplay` is the device address where
  pipeline’s metadata was originally saved and can now be used to
  re-populate `deviceAddress` for replay.

## <a href="#_description" class="anchor"></a>Description

If `pipelineDeviceAddressCaptureReplay` is zero, no specific address is
requested. If `pipelineDeviceAddressCaptureReplay` is not zero, then it
**must** be an address retrieved from an identically created pipeline on
the same implementation. The pipeline metadata **must** also be placed
on an identically created buffer and at the same offset using the
[vkCmdUpdatePipelineIndirectBufferNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdUpdatePipelineIndirectBufferNV.html)
command.

Valid Usage

- <a
  href="#VUID-VkComputePipelineIndirectBufferInfoNV-deviceGeneratedComputePipelines-09009"
  id="VUID-VkComputePipelineIndirectBufferInfoNV-deviceGeneratedComputePipelines-09009"></a>
  VUID-VkComputePipelineIndirectBufferInfoNV-deviceGeneratedComputePipelines-09009  
  The <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-deviceGeneratedComputePipelines"
  target="_blank"
  rel="noopener"><code>VkPhysicalDeviceDeviceGeneratedCommandsComputeFeaturesNV</code>::<code>deviceGeneratedComputePipelines</code></a>
  feature **must** be enabled

- <a href="#VUID-VkComputePipelineIndirectBufferInfoNV-flags-09010"
  id="VUID-VkComputePipelineIndirectBufferInfoNV-flags-09010"></a>
  VUID-VkComputePipelineIndirectBufferInfoNV-flags-09010  
  The pipeline creation flags in
  [VkComputePipelineCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkComputePipelineCreateInfo.html)::`flags`
  **must** include `VK_PIPELINE_CREATE_INDIRECT_BINDABLE_BIT_NV`

- <a
  href="#VUID-VkComputePipelineIndirectBufferInfoNV-deviceAddress-09011"
  id="VUID-VkComputePipelineIndirectBufferInfoNV-deviceAddress-09011"></a>
  VUID-VkComputePipelineIndirectBufferInfoNV-deviceAddress-09011  
  `deviceAddress` **must** be aligned to the
  [VkMemoryRequirements2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryRequirements2.html)::`alignment`, as
  returned by
  [vkGetPipelineIndirectMemoryRequirementsNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPipelineIndirectMemoryRequirementsNV.html)

- <a
  href="#VUID-VkComputePipelineIndirectBufferInfoNV-deviceAddress-09012"
  id="VUID-VkComputePipelineIndirectBufferInfoNV-deviceAddress-09012"></a>
  VUID-VkComputePipelineIndirectBufferInfoNV-deviceAddress-09012  
  `deviceAddress` **must** have been allocated from a buffer that was
  created with usage `VK_BUFFER_USAGE_TRANSFER_DST_BIT` and
  `VK_BUFFER_USAGE_INDIRECT_BUFFER_BIT`

- <a href="#VUID-VkComputePipelineIndirectBufferInfoNV-size-09013"
  id="VUID-VkComputePipelineIndirectBufferInfoNV-size-09013"></a>
  VUID-VkComputePipelineIndirectBufferInfoNV-size-09013  
  `size` **must** be greater than or equal to the
  [VkMemoryRequirements2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryRequirements2.html)::`size`, as
  returned by
  [vkGetPipelineIndirectMemoryRequirementsNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPipelineIndirectMemoryRequirementsNV.html)

- <a
  href="#VUID-VkComputePipelineIndirectBufferInfoNV-pipelineDeviceAddressCaptureReplay-09014"
  id="VUID-VkComputePipelineIndirectBufferInfoNV-pipelineDeviceAddressCaptureReplay-09014"></a>
  VUID-VkComputePipelineIndirectBufferInfoNV-pipelineDeviceAddressCaptureReplay-09014  
  If `pipelineDeviceAddressCaptureReplay` is non-zero then the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-deviceGeneratedComputePipelines"
  target="_blank"
  rel="noopener"><code>VkPhysicalDeviceDeviceGeneratedCommandsComputeFeaturesNV</code>::<code>deviceGeneratedComputeCaptureReplay</code></a>
  feature **must** be enabled

- <a
  href="#VUID-VkComputePipelineIndirectBufferInfoNV-pipelineDeviceAddressCaptureReplay-09015"
  id="VUID-VkComputePipelineIndirectBufferInfoNV-pipelineDeviceAddressCaptureReplay-09015"></a>
  VUID-VkComputePipelineIndirectBufferInfoNV-pipelineDeviceAddressCaptureReplay-09015  
  If `pipelineDeviceAddressCaptureReplay` is non-zero then that address
  **must** have been allocated with flag
  `VK_MEMORY_ALLOCATE_DEVICE_ADDRESS_CAPTURE_REPLAY_BIT` set

- <a
  href="#VUID-VkComputePipelineIndirectBufferInfoNV-pipelineDeviceAddressCaptureReplay-09016"
  id="VUID-VkComputePipelineIndirectBufferInfoNV-pipelineDeviceAddressCaptureReplay-09016"></a>
  VUID-VkComputePipelineIndirectBufferInfoNV-pipelineDeviceAddressCaptureReplay-09016  
  If `pipelineDeviceAddressCaptureReplay` is non-zero, the `pipeline`
  **must** have been recreated for replay

- <a
  href="#VUID-VkComputePipelineIndirectBufferInfoNV-pipelineDeviceAddressCaptureReplay-09017"
  id="VUID-VkComputePipelineIndirectBufferInfoNV-pipelineDeviceAddressCaptureReplay-09017"></a>
  VUID-VkComputePipelineIndirectBufferInfoNV-pipelineDeviceAddressCaptureReplay-09017  
  `pipelineDeviceAddressCaptureReplay` **must** satisfy the `alignment`
  and `size` requirements similar to `deviceAddress`

Valid Usage (Implicit)

- <a href="#VUID-VkComputePipelineIndirectBufferInfoNV-sType-sType"
  id="VUID-VkComputePipelineIndirectBufferInfoNV-sType-sType"></a>
  VUID-VkComputePipelineIndirectBufferInfoNV-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_COMPUTE_PIPELINE_INDIRECT_BUFFER_INFO_NV`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_device_generated_commands_compute](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_device_generated_commands_compute.html),
[VkDeviceAddress](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceAddress.html),
[VkDeviceSize](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceSize.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkComputePipelineIndirectBufferInfoNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
