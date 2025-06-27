# VkIndirectCommandsLayoutCreateInfoNV(3) Manual Page

## Name

VkIndirectCommandsLayoutCreateInfoNV - Structure specifying the
parameters of a newly created indirect commands layout object



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkIndirectCommandsLayoutCreateInfoNV` structure is defined as:

``` c
// Provided by VK_NV_device_generated_commands
typedef struct VkIndirectCommandsLayoutCreateInfoNV {
    VkStructureType                           sType;
    const void*                               pNext;
    VkIndirectCommandsLayoutUsageFlagsNV      flags;
    VkPipelineBindPoint                       pipelineBindPoint;
    uint32_t                                  tokenCount;
    const VkIndirectCommandsLayoutTokenNV*    pTokens;
    uint32_t                                  streamCount;
    const uint32_t*                           pStreamStrides;
} VkIndirectCommandsLayoutCreateInfoNV;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `pipelineBindPoint` is the
  [VkPipelineBindPoint](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineBindPoint.html) that this layout
  targets.

- `flags` is a bitmask of
  [VkIndirectCommandsLayoutUsageFlagBitsNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkIndirectCommandsLayoutUsageFlagBitsNV.html)
  specifying usage hints of this layout.

- `tokenCount` is the length of the individual command sequence.

- `pTokens` is an array describing each command token in detail. See
  [VkIndirectCommandsTokenTypeNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkIndirectCommandsTokenTypeNV.html)
  and
  [VkIndirectCommandsLayoutTokenNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkIndirectCommandsLayoutTokenNV.html)
  below for details.

- `streamCount` is the number of streams used to provide the token
  inputs.

- `pStreamStrides` is an array defining the byte stride for each input
  stream.

## <a href="#_description" class="anchor"></a>Description

The following code illustrates some of the flags:

``` c
void cmdProcessAllSequences(cmd, pipeline, indirectCommandsLayout, pIndirectCommandsTokens, sequencesCount, indexbuffer, indexbufferOffset)
{
  for (s = 0; s < sequencesCount; s++)
  {
    sUsed = s;

    if (indirectCommandsLayout.flags & VK_INDIRECT_COMMANDS_LAYOUT_USAGE_INDEXED_SEQUENCES_BIT_NV) {
      sUsed = indexbuffer.load_uint32( sUsed * sizeof(uint32_t) + indexbufferOffset);
    }

    if (indirectCommandsLayout.flags & VK_INDIRECT_COMMANDS_LAYOUT_USAGE_UNORDERED_SEQUENCES_BIT_NV) {
      sUsed = incoherent_implementation_dependent_permutation[ sUsed ];
    }

    cmdProcessSequence( cmd, pipeline, indirectCommandsLayout, pIndirectCommandsTokens, sUsed );
  }
}
```

When tokens are consumed, an offset is computed based on token offset
and stream stride. The resulting offset is required to be aligned. The
alignment for a specific token is equal to the scalar alignment of the
data type as defined in <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#interfaces-alignment-requirements"
target="_blank" rel="noopener">Alignment Requirements</a>, or
`VkPhysicalDeviceDeviceGeneratedCommandsPropertiesNV`::`minIndirectCommandsBufferOffsetAlignment`,
whichever is lower.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr>
<td class="icon"><em></em></td>
<td class="content">Note
<p>A <code>minIndirectCommandsBufferOffsetAlignment</code> of 4 allows
<a href="VkDeviceAddress.html">VkDeviceAddress</a> to be packed as
<code>uvec2</code> with scalar layout instead of <code>uint64_t</code>
with 8 byte alignment. This enables direct compatibility with D3D12
command signature layouts.</p></td>
</tr>
</tbody>
</table>

Valid Usage

- <a
  href="#VUID-VkIndirectCommandsLayoutCreateInfoNV-pipelineBindPoint-02930"
  id="VUID-VkIndirectCommandsLayoutCreateInfoNV-pipelineBindPoint-02930"></a>
  VUID-VkIndirectCommandsLayoutCreateInfoNV-pipelineBindPoint-02930  
  The `pipelineBindPoint` **must** be `VK_PIPELINE_BIND_POINT_GRAPHICS`
  or `VK_PIPELINE_BIND_POINT_COMPUTE`

- <a href="#VUID-VkIndirectCommandsLayoutCreateInfoNV-tokenCount-02931"
  id="VUID-VkIndirectCommandsLayoutCreateInfoNV-tokenCount-02931"></a>
  VUID-VkIndirectCommandsLayoutCreateInfoNV-tokenCount-02931  
  `tokenCount` **must** be greater than `0` and less than or equal to
  `VkPhysicalDeviceDeviceGeneratedCommandsPropertiesNV`::`maxIndirectCommandsTokenCount`

- <a href="#VUID-VkIndirectCommandsLayoutCreateInfoNV-pTokens-02932"
  id="VUID-VkIndirectCommandsLayoutCreateInfoNV-pTokens-02932"></a>
  VUID-VkIndirectCommandsLayoutCreateInfoNV-pTokens-02932  
  If `pTokens` contains an entry of
  `VK_INDIRECT_COMMANDS_TOKEN_TYPE_SHADER_GROUP_NV` it **must** be the
  first element of the array and there **must** be only a single element
  of such token type

- <a href="#VUID-VkIndirectCommandsLayoutCreateInfoNV-pTokens-09585"
  id="VUID-VkIndirectCommandsLayoutCreateInfoNV-pTokens-09585"></a>
  VUID-VkIndirectCommandsLayoutCreateInfoNV-pTokens-09585  
  If `pTokens` contains an entry of
  `VK_INDIRECT_COMMANDS_TOKEN_TYPE_PIPELINE_NV` it **must** be the first
  element of the array and there **must** be only a single element of
  such token type

- <a href="#VUID-VkIndirectCommandsLayoutCreateInfoNV-pTokens-02933"
  id="VUID-VkIndirectCommandsLayoutCreateInfoNV-pTokens-02933"></a>
  VUID-VkIndirectCommandsLayoutCreateInfoNV-pTokens-02933  
  If `pTokens` contains an entry of
  `VK_INDIRECT_COMMANDS_TOKEN_TYPE_STATE_FLAGS_NV` there **must** be
  only a single element of such token type

- <a href="#VUID-VkIndirectCommandsLayoutCreateInfoNV-pTokens-02934"
  id="VUID-VkIndirectCommandsLayoutCreateInfoNV-pTokens-02934"></a>
  VUID-VkIndirectCommandsLayoutCreateInfoNV-pTokens-02934  
  All state tokens in `pTokens` **must** occur before any action command
  tokens (`VK_INDIRECT_COMMANDS_TOKEN_TYPE_DRAW_NV`,
  `VK_INDIRECT_COMMANDS_TOKEN_TYPE_DRAW_INDEXED_NV`,
  `VK_INDIRECT_COMMANDS_TOKEN_TYPE_DRAW_TASKS_NV`,
  `VK_INDIRECT_COMMANDS_TOKEN_TYPE_DRAW_MESH_TASKS_NV` ,
  `VK_INDIRECT_COMMANDS_TOKEN_TYPE_DISPATCH_NV` )

- <a href="#VUID-VkIndirectCommandsLayoutCreateInfoNV-pTokens-02935"
  id="VUID-VkIndirectCommandsLayoutCreateInfoNV-pTokens-02935"></a>
  VUID-VkIndirectCommandsLayoutCreateInfoNV-pTokens-02935  
  The content of `pTokens` **must** include one single action command
  token that is compatible with the `pipelineBindPoint`

- <a href="#VUID-VkIndirectCommandsLayoutCreateInfoNV-streamCount-02936"
  id="VUID-VkIndirectCommandsLayoutCreateInfoNV-streamCount-02936"></a>
  VUID-VkIndirectCommandsLayoutCreateInfoNV-streamCount-02936  
  `streamCount` **must** be greater than `0` and less or equal to
  `VkPhysicalDeviceDeviceGeneratedCommandsPropertiesNV`::`maxIndirectCommandsStreamCount`

- <a
  href="#VUID-VkIndirectCommandsLayoutCreateInfoNV-pStreamStrides-02937"
  id="VUID-VkIndirectCommandsLayoutCreateInfoNV-pStreamStrides-02937"></a>
  VUID-VkIndirectCommandsLayoutCreateInfoNV-pStreamStrides-02937  
  each element of `pStreamStrides` **must** be greater than `0` and less
  than or equal to
  `VkPhysicalDeviceDeviceGeneratedCommandsPropertiesNV`::`maxIndirectCommandsStreamStride`.
  Furthermore the alignment of each token input **must** be ensured

- <a
  href="#VUID-VkIndirectCommandsLayoutCreateInfoNV-pipelineBindPoint-09088"
  id="VUID-VkIndirectCommandsLayoutCreateInfoNV-pipelineBindPoint-09088"></a>
  VUID-VkIndirectCommandsLayoutCreateInfoNV-pipelineBindPoint-09088  
  If `pipelineBindPoint` is `VK_PIPELINE_BIND_POINT_COMPUTE` then the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-deviceGeneratedCompute"
  target="_blank"
  rel="noopener"><code>VkPhysicalDeviceDeviceGeneratedCommandsComputeFeaturesNV</code>::<code>deviceGeneratedCompute</code></a>
  feature **must** be enabled

- <a
  href="#VUID-VkIndirectCommandsLayoutCreateInfoNV-pipelineBindPoint-09089"
  id="VUID-VkIndirectCommandsLayoutCreateInfoNV-pipelineBindPoint-09089"></a>
  VUID-VkIndirectCommandsLayoutCreateInfoNV-pipelineBindPoint-09089  
  If `pipelineBindPoint` is `VK_PIPELINE_BIND_POINT_COMPUTE` then the
  state tokens in `pTokens` **must** only include
  `VK_INDIRECT_COMMANDS_TOKEN_TYPE_DISPATCH_NV`,
  `VK_INDIRECT_COMMANDS_TOKEN_TYPE_PIPELINE_NV`, or
  `VK_INDIRECT_COMMANDS_TOKEN_TYPE_PUSH_CONSTANT_NV`

- <a
  href="#VUID-VkIndirectCommandsLayoutCreateInfoNV-pipelineBindPoint-09090"
  id="VUID-VkIndirectCommandsLayoutCreateInfoNV-pipelineBindPoint-09090"></a>
  VUID-VkIndirectCommandsLayoutCreateInfoNV-pipelineBindPoint-09090  
  If `pipelineBindPoint` is `VK_PIPELINE_BIND_POINT_COMPUTE` and
  `pTokens` includes `VK_INDIRECT_COMMANDS_TOKEN_TYPE_PIPELINE_NV`, then
  the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-deviceGeneratedComputePipelines"
  target="_blank"
  rel="noopener"><code>VkPhysicalDeviceDeviceGeneratedCommandsComputeFeaturesNV</code>::<code>deviceGeneratedComputePipelines</code></a>
  feature **must** be enabled

Valid Usage (Implicit)

- <a href="#VUID-VkIndirectCommandsLayoutCreateInfoNV-sType-sType"
  id="VUID-VkIndirectCommandsLayoutCreateInfoNV-sType-sType"></a>
  VUID-VkIndirectCommandsLayoutCreateInfoNV-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_INDIRECT_COMMANDS_LAYOUT_CREATE_INFO_NV`

- <a href="#VUID-VkIndirectCommandsLayoutCreateInfoNV-pNext-pNext"
  id="VUID-VkIndirectCommandsLayoutCreateInfoNV-pNext-pNext"></a>
  VUID-VkIndirectCommandsLayoutCreateInfoNV-pNext-pNext  
  `pNext` **must** be `NULL`

- <a href="#VUID-VkIndirectCommandsLayoutCreateInfoNV-flags-parameter"
  id="VUID-VkIndirectCommandsLayoutCreateInfoNV-flags-parameter"></a>
  VUID-VkIndirectCommandsLayoutCreateInfoNV-flags-parameter  
  `flags` **must** be a valid combination of
  [VkIndirectCommandsLayoutUsageFlagBitsNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkIndirectCommandsLayoutUsageFlagBitsNV.html)
  values

- <a
  href="#VUID-VkIndirectCommandsLayoutCreateInfoNV-pipelineBindPoint-parameter"
  id="VUID-VkIndirectCommandsLayoutCreateInfoNV-pipelineBindPoint-parameter"></a>
  VUID-VkIndirectCommandsLayoutCreateInfoNV-pipelineBindPoint-parameter  
  `pipelineBindPoint` **must** be a valid
  [VkPipelineBindPoint](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineBindPoint.html) value

- <a href="#VUID-VkIndirectCommandsLayoutCreateInfoNV-pTokens-parameter"
  id="VUID-VkIndirectCommandsLayoutCreateInfoNV-pTokens-parameter"></a>
  VUID-VkIndirectCommandsLayoutCreateInfoNV-pTokens-parameter  
  `pTokens` **must** be a valid pointer to an array of `tokenCount`
  valid
  [VkIndirectCommandsLayoutTokenNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkIndirectCommandsLayoutTokenNV.html)
  structures

- <a
  href="#VUID-VkIndirectCommandsLayoutCreateInfoNV-pStreamStrides-parameter"
  id="VUID-VkIndirectCommandsLayoutCreateInfoNV-pStreamStrides-parameter"></a>
  VUID-VkIndirectCommandsLayoutCreateInfoNV-pStreamStrides-parameter  
  `pStreamStrides` **must** be a valid pointer to an array of
  `streamCount` `uint32_t` values

- <a
  href="#VUID-VkIndirectCommandsLayoutCreateInfoNV-tokenCount-arraylength"
  id="VUID-VkIndirectCommandsLayoutCreateInfoNV-tokenCount-arraylength"></a>
  VUID-VkIndirectCommandsLayoutCreateInfoNV-tokenCount-arraylength  
  `tokenCount` **must** be greater than `0`

- <a
  href="#VUID-VkIndirectCommandsLayoutCreateInfoNV-streamCount-arraylength"
  id="VUID-VkIndirectCommandsLayoutCreateInfoNV-streamCount-arraylength"></a>
  VUID-VkIndirectCommandsLayoutCreateInfoNV-streamCount-arraylength  
  `streamCount` **must** be greater than `0`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_device_generated_commands](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_device_generated_commands.html),
[VkIndirectCommandsLayoutTokenNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkIndirectCommandsLayoutTokenNV.html),
[VkIndirectCommandsLayoutUsageFlagsNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkIndirectCommandsLayoutUsageFlagsNV.html),
[VkPipelineBindPoint](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineBindPoint.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkCreateIndirectCommandsLayoutNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateIndirectCommandsLayoutNV.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkIndirectCommandsLayoutCreateInfoNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
