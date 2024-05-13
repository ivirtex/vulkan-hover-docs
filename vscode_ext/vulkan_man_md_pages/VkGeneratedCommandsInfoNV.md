# VkGeneratedCommandsInfoNV(3) Manual Page

## Name

VkGeneratedCommandsInfoNV - Structure specifying parameters for the
generation of commands



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkGeneratedCommandsInfoNV` is defined as:

``` c
// Provided by VK_NV_device_generated_commands
typedef struct VkGeneratedCommandsInfoNV {
    VkStructureType                      sType;
    const void*                          pNext;
    VkPipelineBindPoint                  pipelineBindPoint;
    VkPipeline                           pipeline;
    VkIndirectCommandsLayoutNV           indirectCommandsLayout;
    uint32_t                             streamCount;
    const VkIndirectCommandsStreamNV*    pStreams;
    uint32_t                             sequencesCount;
    VkBuffer                             preprocessBuffer;
    VkDeviceSize                         preprocessOffset;
    VkDeviceSize                         preprocessSize;
    VkBuffer                             sequencesCountBuffer;
    VkDeviceSize                         sequencesCountOffset;
    VkBuffer                             sequencesIndexBuffer;
    VkDeviceSize                         sequencesIndexOffset;
} VkGeneratedCommandsInfoNV;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `pipelineBindPoint` is the
  [VkPipelineBindPoint](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineBindPoint.html) used for the
  `pipeline`.

- `pipeline` is the [VkPipeline](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipeline.html) used in the generation
  and execution process.

- `indirectCommandsLayout` is the
  [VkIndirectCommandsLayoutNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkIndirectCommandsLayoutNV.html) that
  provides the command sequence to generate.

- `streamCount` defines the number of input streams

- `pStreams` is a pointer to an array of `streamCount`
  [VkIndirectCommandsStreamNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkIndirectCommandsStreamNV.html)
  structures providing the input data for the tokens used in
  `indirectCommandsLayout`.

- `sequencesCount` is the maximum number of sequences to reserve. If
  `sequencesCountBuffer` is [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), this
  is also the actual number of sequences generated.

- `preprocessBuffer` is the [VkBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBuffer.html) that is used for
  preprocessing the input data for execution. If this structure is used
  with
  [vkCmdExecuteGeneratedCommandsNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdExecuteGeneratedCommandsNV.html)
  with its `isPreprocessed` set to `VK_TRUE`, then the preprocessing
  step is skipped and data in this buffer will not be modified. The
  contents and the layout of this buffer are opaque to applications and
  **must** not be modified outside functions related to device-generated
  commands or copied to another buffer for reuse.

- `preprocessOffset` is the byte offset into `preprocessBuffer` where
  the preprocessed data is stored.

- `preprocessSize` is the maximum byte size within the
  `preprocessBuffer` after the `preprocessOffset` that is available for
  preprocessing.

- `sequencesCountBuffer` is a `VkBuffer` in which the actual number of
  sequences is provided as single `uint32_t` value.

- `sequencesCountOffset` is the byte offset into `sequencesCountBuffer`
  where the count value is stored.

- `sequencesIndexBuffer` is a `VkBuffer` that encodes the used sequence
  indices as `uint32_t` array.

- `sequencesIndexOffset` is the byte offset into `sequencesIndexBuffer`
  where the index values start.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-VkGeneratedCommandsInfoNV-pipeline-02912"
  id="VUID-VkGeneratedCommandsInfoNV-pipeline-02912"></a>
  VUID-VkGeneratedCommandsInfoNV-pipeline-02912  
  The provided `pipeline` **must** match the pipeline bound at execution
  time

- <a href="#VUID-VkGeneratedCommandsInfoNV-indirectCommandsLayout-02913"
  id="VUID-VkGeneratedCommandsInfoNV-indirectCommandsLayout-02913"></a>
  VUID-VkGeneratedCommandsInfoNV-indirectCommandsLayout-02913  
  If the `indirectCommandsLayout` uses a token of
  `VK_INDIRECT_COMMANDS_TOKEN_TYPE_SHADER_GROUP_NV`, then the `pipeline`
  **must** have been created with multiple shader groups

- <a href="#VUID-VkGeneratedCommandsInfoNV-indirectCommandsLayout-02914"
  id="VUID-VkGeneratedCommandsInfoNV-indirectCommandsLayout-02914"></a>
  VUID-VkGeneratedCommandsInfoNV-indirectCommandsLayout-02914  
  If the `indirectCommandsLayout` uses a token of
  `VK_INDIRECT_COMMANDS_TOKEN_TYPE_SHADER_GROUP_NV`, then the `pipeline`
  **must** have been created with
  `VK_PIPELINE_CREATE_INDIRECT_BINDABLE_BIT_NV` set in
  `VkGraphicsPipelineCreateInfo`::`flags`

- <a href="#VUID-VkGeneratedCommandsInfoNV-indirectCommandsLayout-02915"
  id="VUID-VkGeneratedCommandsInfoNV-indirectCommandsLayout-02915"></a>
  VUID-VkGeneratedCommandsInfoNV-indirectCommandsLayout-02915  
  If the `indirectCommandsLayout` uses a token of
  `VK_INDIRECT_COMMANDS_TOKEN_TYPE_PUSH_CONSTANT_NV`, then the
  `pipeline`’s `VkPipelineLayout` **must** match the
  [VkIndirectCommandsLayoutTokenNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkIndirectCommandsLayoutTokenNV.html)::`pushconstantPipelineLayout`

- <a href="#VUID-VkGeneratedCommandsInfoNV-streamCount-02916"
  id="VUID-VkGeneratedCommandsInfoNV-streamCount-02916"></a>
  VUID-VkGeneratedCommandsInfoNV-streamCount-02916  
  `streamCount` **must** match the `indirectCommandsLayout`’s
  `streamCount`

- <a href="#VUID-VkGeneratedCommandsInfoNV-pipelineBindPoint-09084"
  id="VUID-VkGeneratedCommandsInfoNV-pipelineBindPoint-09084"></a>
  VUID-VkGeneratedCommandsInfoNV-pipelineBindPoint-09084  
  If `pipelineBindPoint` is of type `VK_PIPELINE_BIND_POINT_COMPUTE`,
  then the `pipeline` **must** have been created with the flag
  `VK_PIPELINE_CREATE_INDIRECT_BINDABLE_BIT_NV`

- <a href="#VUID-VkGeneratedCommandsInfoNV-pipelineBindPoint-09085"
  id="VUID-VkGeneratedCommandsInfoNV-pipelineBindPoint-09085"></a>
  VUID-VkGeneratedCommandsInfoNV-pipelineBindPoint-09085  
  If `pipelineBindPoint` is of type `VK_PIPELINE_BIND_POINT_COMPUTE`,
  then the `pipeline` **must** have been created with a
  [VkComputePipelineIndirectBufferInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkComputePipelineIndirectBufferInfoNV.html)
  structure specifying a valid address where its metadata will be saved

- <a href="#VUID-VkGeneratedCommandsInfoNV-pipelineBindPoint-09086"
  id="VUID-VkGeneratedCommandsInfoNV-pipelineBindPoint-09086"></a>
  VUID-VkGeneratedCommandsInfoNV-pipelineBindPoint-09086  
  If `pipelineBindPoint` is of type `VK_PIPELINE_BIND_POINT_COMPUTE`,
  then
  [vkCmdUpdatePipelineIndirectBufferNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdUpdatePipelineIndirectBufferNV.html)
  **must** have been called on that pipeline to save its metadata to a
  device address

- <a href="#VUID-VkGeneratedCommandsInfoNV-pipelineBindPoint-09087"
  id="VUID-VkGeneratedCommandsInfoNV-pipelineBindPoint-09087"></a>
  VUID-VkGeneratedCommandsInfoNV-pipelineBindPoint-09087  
  If `pipelineBindPoint` is of type `VK_PIPELINE_BIND_POINT_COMPUTE`,
  and if `VK_INDIRECT_COMMANDS_TOKEN_TYPE_PIPELINE_NV` is used, then
  `pipeline` **must** be [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html)

- <a href="#VUID-VkGeneratedCommandsInfoNV-sequencesCount-02917"
  id="VUID-VkGeneratedCommandsInfoNV-sequencesCount-02917"></a>
  VUID-VkGeneratedCommandsInfoNV-sequencesCount-02917  
  `sequencesCount` **must** be less or equal to
  [VkPhysicalDeviceDeviceGeneratedCommandsPropertiesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceDeviceGeneratedCommandsPropertiesNV.html)::`maxIndirectSequenceCount`
  and
  [VkGeneratedCommandsMemoryRequirementsInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGeneratedCommandsMemoryRequirementsInfoNV.html)::`maxSequencesCount`
  that was used to determine the `preprocessSize`

- <a href="#VUID-VkGeneratedCommandsInfoNV-preprocessBuffer-02918"
  id="VUID-VkGeneratedCommandsInfoNV-preprocessBuffer-02918"></a>
  VUID-VkGeneratedCommandsInfoNV-preprocessBuffer-02918  
  `preprocessBuffer` **must** have the
  `VK_BUFFER_USAGE_INDIRECT_BUFFER_BIT` bit set in its usage flag

- <a href="#VUID-VkGeneratedCommandsInfoNV-preprocessOffset-02919"
  id="VUID-VkGeneratedCommandsInfoNV-preprocessOffset-02919"></a>
  VUID-VkGeneratedCommandsInfoNV-preprocessOffset-02919  
  `preprocessOffset` **must** be aligned to
  [VkPhysicalDeviceDeviceGeneratedCommandsPropertiesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceDeviceGeneratedCommandsPropertiesNV.html)::`minIndirectCommandsBufferOffsetAlignment`

- <a href="#VUID-VkGeneratedCommandsInfoNV-preprocessBuffer-02971"
  id="VUID-VkGeneratedCommandsInfoNV-preprocessBuffer-02971"></a>
  VUID-VkGeneratedCommandsInfoNV-preprocessBuffer-02971  
  If `preprocessBuffer` is non-sparse then it **must** be bound
  completely and contiguously to a single `VkDeviceMemory` object

- <a href="#VUID-VkGeneratedCommandsInfoNV-preprocessSize-02920"
  id="VUID-VkGeneratedCommandsInfoNV-preprocessSize-02920"></a>
  VUID-VkGeneratedCommandsInfoNV-preprocessSize-02920  
  `preprocessSize` **must** be at least equal to the memory
  requirement\`s size returned by
  [vkGetGeneratedCommandsMemoryRequirementsNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetGeneratedCommandsMemoryRequirementsNV.html)
  using the matching inputs (`indirectCommandsLayout`, …​) as within this
  structure

- <a href="#VUID-VkGeneratedCommandsInfoNV-sequencesCountBuffer-02921"
  id="VUID-VkGeneratedCommandsInfoNV-sequencesCountBuffer-02921"></a>
  VUID-VkGeneratedCommandsInfoNV-sequencesCountBuffer-02921  
  `sequencesCountBuffer` **can** be set if the actual used count of
  sequences is sourced from the provided buffer. In that case the
  `sequencesCount` serves as upper bound

- <a href="#VUID-VkGeneratedCommandsInfoNV-sequencesCountBuffer-02922"
  id="VUID-VkGeneratedCommandsInfoNV-sequencesCountBuffer-02922"></a>
  VUID-VkGeneratedCommandsInfoNV-sequencesCountBuffer-02922  
  If `sequencesCountBuffer` is not
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), its usage flag **must** have
  the `VK_BUFFER_USAGE_INDIRECT_BUFFER_BIT` bit set

- <a href="#VUID-VkGeneratedCommandsInfoNV-sequencesCountBuffer-02923"
  id="VUID-VkGeneratedCommandsInfoNV-sequencesCountBuffer-02923"></a>
  VUID-VkGeneratedCommandsInfoNV-sequencesCountBuffer-02923  
  If `sequencesCountBuffer` is not
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), `sequencesCountOffset` **must**
  be aligned to
  `VkPhysicalDeviceDeviceGeneratedCommandsPropertiesNV`::`minSequencesCountBufferOffsetAlignment`

- <a href="#VUID-VkGeneratedCommandsInfoNV-sequencesCountBuffer-02972"
  id="VUID-VkGeneratedCommandsInfoNV-sequencesCountBuffer-02972"></a>
  VUID-VkGeneratedCommandsInfoNV-sequencesCountBuffer-02972  
  If `sequencesCountBuffer` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html)
  and is non-sparse then it **must** be bound completely and
  contiguously to a single `VkDeviceMemory` object

- <a href="#VUID-VkGeneratedCommandsInfoNV-sequencesIndexBuffer-02924"
  id="VUID-VkGeneratedCommandsInfoNV-sequencesIndexBuffer-02924"></a>
  VUID-VkGeneratedCommandsInfoNV-sequencesIndexBuffer-02924  
  If `indirectCommandsLayout`’s
  `VK_INDIRECT_COMMANDS_LAYOUT_USAGE_INDEXED_SEQUENCES_BIT_NV` is set,
  `sequencesIndexBuffer` **must** be set otherwise it **must** be
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html)

- <a href="#VUID-VkGeneratedCommandsInfoNV-sequencesIndexBuffer-02925"
  id="VUID-VkGeneratedCommandsInfoNV-sequencesIndexBuffer-02925"></a>
  VUID-VkGeneratedCommandsInfoNV-sequencesIndexBuffer-02925  
  If `sequencesIndexBuffer` is not
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), its usage flag **must** have
  the `VK_BUFFER_USAGE_INDIRECT_BUFFER_BIT` bit set

- <a href="#VUID-VkGeneratedCommandsInfoNV-sequencesIndexBuffer-02926"
  id="VUID-VkGeneratedCommandsInfoNV-sequencesIndexBuffer-02926"></a>
  VUID-VkGeneratedCommandsInfoNV-sequencesIndexBuffer-02926  
  If `sequencesIndexBuffer` is not
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), `sequencesIndexOffset` **must**
  be aligned to
  `VkPhysicalDeviceDeviceGeneratedCommandsPropertiesNV`::`minSequencesIndexBufferOffsetAlignment`

- <a href="#VUID-VkGeneratedCommandsInfoNV-sequencesIndexBuffer-02973"
  id="VUID-VkGeneratedCommandsInfoNV-sequencesIndexBuffer-02973"></a>
  VUID-VkGeneratedCommandsInfoNV-sequencesIndexBuffer-02973  
  If `sequencesIndexBuffer` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html)
  and is non-sparse then it **must** be bound completely and
  contiguously to a single `VkDeviceMemory` object

- <a href="#VUID-VkGeneratedCommandsInfoNV-indirectCommandsLayout-07078"
  id="VUID-VkGeneratedCommandsInfoNV-indirectCommandsLayout-07078"></a>
  VUID-VkGeneratedCommandsInfoNV-indirectCommandsLayout-07078  
  If the `indirectCommandsLayout` uses a token of
  `VK_INDIRECT_COMMANDS_TOKEN_TYPE_DRAW_TASKS_NV`, then the `pipeline`
  **must** contain a shader stage using the `MeshNV` `Execution` `Model`

- <a href="#VUID-VkGeneratedCommandsInfoNV-indirectCommandsLayout-07079"
  id="VUID-VkGeneratedCommandsInfoNV-indirectCommandsLayout-07079"></a>
  VUID-VkGeneratedCommandsInfoNV-indirectCommandsLayout-07079  
  If the `indirectCommandsLayout` uses a token of
  `VK_INDIRECT_COMMANDS_TOKEN_TYPE_DRAW_MESH_TASKS_NV`, then the
  `pipeline` **must** contain a shader stage using the `MeshEXT`
  `Execution` `Model`

Valid Usage (Implicit)

- <a href="#VUID-VkGeneratedCommandsInfoNV-sType-sType"
  id="VUID-VkGeneratedCommandsInfoNV-sType-sType"></a>
  VUID-VkGeneratedCommandsInfoNV-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_GENERATED_COMMANDS_INFO_NV`

- <a href="#VUID-VkGeneratedCommandsInfoNV-pNext-pNext"
  id="VUID-VkGeneratedCommandsInfoNV-pNext-pNext"></a>
  VUID-VkGeneratedCommandsInfoNV-pNext-pNext  
  `pNext` **must** be `NULL`

- <a href="#VUID-VkGeneratedCommandsInfoNV-pipelineBindPoint-parameter"
  id="VUID-VkGeneratedCommandsInfoNV-pipelineBindPoint-parameter"></a>
  VUID-VkGeneratedCommandsInfoNV-pipelineBindPoint-parameter  
  `pipelineBindPoint` **must** be a valid
  [VkPipelineBindPoint](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineBindPoint.html) value

- <a href="#VUID-VkGeneratedCommandsInfoNV-pipeline-parameter"
  id="VUID-VkGeneratedCommandsInfoNV-pipeline-parameter"></a>
  VUID-VkGeneratedCommandsInfoNV-pipeline-parameter  
  If `pipeline` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), `pipeline`
  **must** be a valid [VkPipeline](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipeline.html) handle

- <a
  href="#VUID-VkGeneratedCommandsInfoNV-indirectCommandsLayout-parameter"
  id="VUID-VkGeneratedCommandsInfoNV-indirectCommandsLayout-parameter"></a>
  VUID-VkGeneratedCommandsInfoNV-indirectCommandsLayout-parameter  
  `indirectCommandsLayout` **must** be a valid
  [VkIndirectCommandsLayoutNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkIndirectCommandsLayoutNV.html) handle

- <a href="#VUID-VkGeneratedCommandsInfoNV-pStreams-parameter"
  id="VUID-VkGeneratedCommandsInfoNV-pStreams-parameter"></a>
  VUID-VkGeneratedCommandsInfoNV-pStreams-parameter  
  `pStreams` **must** be a valid pointer to an array of `streamCount`
  valid [VkIndirectCommandsStreamNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkIndirectCommandsStreamNV.html)
  structures

- <a href="#VUID-VkGeneratedCommandsInfoNV-preprocessBuffer-parameter"
  id="VUID-VkGeneratedCommandsInfoNV-preprocessBuffer-parameter"></a>
  VUID-VkGeneratedCommandsInfoNV-preprocessBuffer-parameter  
  `preprocessBuffer` **must** be a valid [VkBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBuffer.html)
  handle

- <a href="#VUID-VkGeneratedCommandsInfoNV-sequencesCountBuffer-parameter"
  id="VUID-VkGeneratedCommandsInfoNV-sequencesCountBuffer-parameter"></a>
  VUID-VkGeneratedCommandsInfoNV-sequencesCountBuffer-parameter  
  If `sequencesCountBuffer` is not
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), `sequencesCountBuffer` **must**
  be a valid [VkBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBuffer.html) handle

- <a href="#VUID-VkGeneratedCommandsInfoNV-sequencesIndexBuffer-parameter"
  id="VUID-VkGeneratedCommandsInfoNV-sequencesIndexBuffer-parameter"></a>
  VUID-VkGeneratedCommandsInfoNV-sequencesIndexBuffer-parameter  
  If `sequencesIndexBuffer` is not
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), `sequencesIndexBuffer` **must**
  be a valid [VkBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBuffer.html) handle

- <a href="#VUID-VkGeneratedCommandsInfoNV-streamCount-arraylength"
  id="VUID-VkGeneratedCommandsInfoNV-streamCount-arraylength"></a>
  VUID-VkGeneratedCommandsInfoNV-streamCount-arraylength  
  `streamCount` **must** be greater than `0`

- <a href="#VUID-VkGeneratedCommandsInfoNV-commonparent"
  id="VUID-VkGeneratedCommandsInfoNV-commonparent"></a>
  VUID-VkGeneratedCommandsInfoNV-commonparent  
  Each of `indirectCommandsLayout`, `pipeline`, `preprocessBuffer`,
  `sequencesCountBuffer`, and `sequencesIndexBuffer` that are valid
  handles of non-ignored parameters **must** have been created,
  allocated, or retrieved from the same [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html)

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_device_generated_commands](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_device_generated_commands.html),
[VkBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBuffer.html), [VkDeviceSize](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceSize.html),
[VkIndirectCommandsLayoutNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkIndirectCommandsLayoutNV.html),
[VkIndirectCommandsStreamNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkIndirectCommandsStreamNV.html),
[VkPipeline](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipeline.html),
[VkPipelineBindPoint](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineBindPoint.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkCmdExecuteGeneratedCommandsNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdExecuteGeneratedCommandsNV.html),
[vkCmdPreprocessGeneratedCommandsNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdPreprocessGeneratedCommandsNV.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkGeneratedCommandsInfoNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
