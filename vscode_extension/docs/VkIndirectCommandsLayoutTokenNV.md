# VkIndirectCommandsLayoutTokenNV(3) Manual Page

## Name

VkIndirectCommandsLayoutTokenNV - Struct specifying the details of an
indirect command layout token



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkIndirectCommandsLayoutTokenNV` structure specifies details to the
function arguments that need to be known at layout creation time:

``` c
// Provided by VK_NV_device_generated_commands
typedef struct VkIndirectCommandsLayoutTokenNV {
    VkStructureType                  sType;
    const void*                      pNext;
    VkIndirectCommandsTokenTypeNV    tokenType;
    uint32_t                         stream;
    uint32_t                         offset;
    uint32_t                         vertexBindingUnit;
    VkBool32                         vertexDynamicStride;
    VkPipelineLayout                 pushconstantPipelineLayout;
    VkShaderStageFlags               pushconstantShaderStageFlags;
    uint32_t                         pushconstantOffset;
    uint32_t                         pushconstantSize;
    VkIndirectStateFlagsNV           indirectStateFlags;
    uint32_t                         indexTypeCount;
    const VkIndexType*               pIndexTypes;
    const uint32_t*                  pIndexTypeValues;
} VkIndirectCommandsLayoutTokenNV;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `tokenType` is a
  [VkIndirectCommandsTokenTypeNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkIndirectCommandsTokenTypeNV.html)
  specifying the token command type.

- `stream` is the index of the input stream containing the token
  argument data.

- `offset` is a relative starting offset within the input stream memory
  for the token argument data.

- `vertexBindingUnit` is used for the vertex buffer binding command.

- `vertexDynamicStride` sets if the vertex buffer stride is provided by
  the binding command rather than the current bound graphics pipeline
  state.

- `pushconstantPipelineLayout` is the `VkPipelineLayout` used for the
  push constant command.

- `pushconstantShaderStageFlags` are the shader stage flags used for the
  push constant command.

- `pushconstantOffset` is the offset used for the push constant command.

- `pushconstantSize` is the size used for the push constant command.

- `indirectStateFlags` is a
  [VkIndirectStateFlagsNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkIndirectStateFlagsNV.html) bitfield
  indicating the active states for the state flag command.

- `indexTypeCount` is the optional size of the `pIndexTypes` and
  `pIndexTypeValues` array pairings. If not zero, it allows to register
  a custom `uint32_t` value to be treated as specific
  [VkIndexType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkIndexType.html).

- `pIndexTypes` is the used [VkIndexType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkIndexType.html) for the
  corresponding `uint32_t` value entry in `pIndexTypeValues`.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-VkIndirectCommandsLayoutTokenNV-stream-02951"
  id="VUID-VkIndirectCommandsLayoutTokenNV-stream-02951"></a>
  VUID-VkIndirectCommandsLayoutTokenNV-stream-02951  
  `stream` **must** be smaller than
  `VkIndirectCommandsLayoutCreateInfoNV`::`streamCount`

- <a href="#VUID-VkIndirectCommandsLayoutTokenNV-offset-02952"
  id="VUID-VkIndirectCommandsLayoutTokenNV-offset-02952"></a>
  VUID-VkIndirectCommandsLayoutTokenNV-offset-02952  
  `offset` **must** be less than or equal to
  `VkPhysicalDeviceDeviceGeneratedCommandsPropertiesNV`::`maxIndirectCommandsTokenOffset`

- <a href="#VUID-VkIndirectCommandsLayoutTokenNV-offset-06888"
  id="VUID-VkIndirectCommandsLayoutTokenNV-offset-06888"></a>
  VUID-VkIndirectCommandsLayoutTokenNV-offset-06888  
  `offset` **must** be aligned to the scalar alignment of `tokenType` or
  `minIndirectCommandsBufferOffsetAlignment`, whichever is lower

- <a href="#VUID-VkIndirectCommandsLayoutTokenNV-tokenType-02976"
  id="VUID-VkIndirectCommandsLayoutTokenNV-tokenType-02976"></a>
  VUID-VkIndirectCommandsLayoutTokenNV-tokenType-02976  
  If `tokenType` is `VK_INDIRECT_COMMANDS_TOKEN_TYPE_VERTEX_BUFFER_NV`,
  `vertexBindingUnit` **must** stay within device supported limits for
  the appropriate commands

- <a href="#VUID-VkIndirectCommandsLayoutTokenNV-tokenType-02977"
  id="VUID-VkIndirectCommandsLayoutTokenNV-tokenType-02977"></a>
  VUID-VkIndirectCommandsLayoutTokenNV-tokenType-02977  
  If `tokenType` is `VK_INDIRECT_COMMANDS_TOKEN_TYPE_PUSH_CONSTANT_NV`,
  `pushconstantPipelineLayout` **must** be valid

- <a href="#VUID-VkIndirectCommandsLayoutTokenNV-tokenType-02978"
  id="VUID-VkIndirectCommandsLayoutTokenNV-tokenType-02978"></a>
  VUID-VkIndirectCommandsLayoutTokenNV-tokenType-02978  
  If `tokenType` is `VK_INDIRECT_COMMANDS_TOKEN_TYPE_PUSH_CONSTANT_NV`,
  `pushconstantOffset` **must** be a multiple of `4`

- <a href="#VUID-VkIndirectCommandsLayoutTokenNV-tokenType-02979"
  id="VUID-VkIndirectCommandsLayoutTokenNV-tokenType-02979"></a>
  VUID-VkIndirectCommandsLayoutTokenNV-tokenType-02979  
  If `tokenType` is `VK_INDIRECT_COMMANDS_TOKEN_TYPE_PUSH_CONSTANT_NV`,
  `pushconstantSize` **must** be a multiple of `4`

- <a href="#VUID-VkIndirectCommandsLayoutTokenNV-tokenType-02980"
  id="VUID-VkIndirectCommandsLayoutTokenNV-tokenType-02980"></a>
  VUID-VkIndirectCommandsLayoutTokenNV-tokenType-02980  
  If `tokenType` is `VK_INDIRECT_COMMANDS_TOKEN_TYPE_PUSH_CONSTANT_NV`,
  `pushconstantOffset` **must** be less than
  `VkPhysicalDeviceLimits`::`maxPushConstantsSize`

- <a href="#VUID-VkIndirectCommandsLayoutTokenNV-tokenType-02981"
  id="VUID-VkIndirectCommandsLayoutTokenNV-tokenType-02981"></a>
  VUID-VkIndirectCommandsLayoutTokenNV-tokenType-02981  
  If `tokenType` is `VK_INDIRECT_COMMANDS_TOKEN_TYPE_PUSH_CONSTANT_NV`,
  `pushconstantSize` **must** be less than or equal to
  `VkPhysicalDeviceLimits`::`maxPushConstantsSize` minus
  `pushconstantOffset`

- <a href="#VUID-VkIndirectCommandsLayoutTokenNV-tokenType-02982"
  id="VUID-VkIndirectCommandsLayoutTokenNV-tokenType-02982"></a>
  VUID-VkIndirectCommandsLayoutTokenNV-tokenType-02982  
  If `tokenType` is `VK_INDIRECT_COMMANDS_TOKEN_TYPE_PUSH_CONSTANT_NV`,
  for each byte in the range specified by `pushconstantOffset` and
  `pushconstantSize` and for each shader stage in
  `pushconstantShaderStageFlags`, there **must** be a push constant
  range in `pushconstantPipelineLayout` that includes that byte and that
  stage

- <a href="#VUID-VkIndirectCommandsLayoutTokenNV-tokenType-02983"
  id="VUID-VkIndirectCommandsLayoutTokenNV-tokenType-02983"></a>
  VUID-VkIndirectCommandsLayoutTokenNV-tokenType-02983  
  If `tokenType` is `VK_INDIRECT_COMMANDS_TOKEN_TYPE_PUSH_CONSTANT_NV`,
  for each byte in the range specified by `pushconstantOffset` and
  `pushconstantSize` and for each push constant range that overlaps that
  byte, `pushconstantShaderStageFlags` **must** include all stages in
  that push constant range’s
  [VkPushConstantRange](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPushConstantRange.html)::`stageFlags`

- <a href="#VUID-VkIndirectCommandsLayoutTokenNV-tokenType-02984"
  id="VUID-VkIndirectCommandsLayoutTokenNV-tokenType-02984"></a>
  VUID-VkIndirectCommandsLayoutTokenNV-tokenType-02984  
  If `tokenType` is `VK_INDIRECT_COMMANDS_TOKEN_TYPE_STATE_FLAGS_NV`,
  `indirectStateFlags` **must** not be `0`

Valid Usage (Implicit)

- <a href="#VUID-VkIndirectCommandsLayoutTokenNV-sType-sType"
  id="VUID-VkIndirectCommandsLayoutTokenNV-sType-sType"></a>
  VUID-VkIndirectCommandsLayoutTokenNV-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_INDIRECT_COMMANDS_LAYOUT_TOKEN_NV`

- <a href="#VUID-VkIndirectCommandsLayoutTokenNV-pNext-pNext"
  id="VUID-VkIndirectCommandsLayoutTokenNV-pNext-pNext"></a>
  VUID-VkIndirectCommandsLayoutTokenNV-pNext-pNext  
  `pNext` **must** be `NULL`

- <a href="#VUID-VkIndirectCommandsLayoutTokenNV-tokenType-parameter"
  id="VUID-VkIndirectCommandsLayoutTokenNV-tokenType-parameter"></a>
  VUID-VkIndirectCommandsLayoutTokenNV-tokenType-parameter  
  `tokenType` **must** be a valid
  [VkIndirectCommandsTokenTypeNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkIndirectCommandsTokenTypeNV.html)
  value

- <a
  href="#VUID-VkIndirectCommandsLayoutTokenNV-pushconstantPipelineLayout-parameter"
  id="VUID-VkIndirectCommandsLayoutTokenNV-pushconstantPipelineLayout-parameter"></a>
  VUID-VkIndirectCommandsLayoutTokenNV-pushconstantPipelineLayout-parameter  
  If `pushconstantPipelineLayout` is not
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), `pushconstantPipelineLayout`
  **must** be a valid [VkPipelineLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLayout.html) handle

- <a
  href="#VUID-VkIndirectCommandsLayoutTokenNV-pushconstantShaderStageFlags-parameter"
  id="VUID-VkIndirectCommandsLayoutTokenNV-pushconstantShaderStageFlags-parameter"></a>
  VUID-VkIndirectCommandsLayoutTokenNV-pushconstantShaderStageFlags-parameter  
  `pushconstantShaderStageFlags` **must** be a valid combination of
  [VkShaderStageFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderStageFlagBits.html) values

- <a
  href="#VUID-VkIndirectCommandsLayoutTokenNV-indirectStateFlags-parameter"
  id="VUID-VkIndirectCommandsLayoutTokenNV-indirectStateFlags-parameter"></a>
  VUID-VkIndirectCommandsLayoutTokenNV-indirectStateFlags-parameter  
  `indirectStateFlags` **must** be a valid combination of
  [VkIndirectStateFlagBitsNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkIndirectStateFlagBitsNV.html) values

- <a href="#VUID-VkIndirectCommandsLayoutTokenNV-pIndexTypes-parameter"
  id="VUID-VkIndirectCommandsLayoutTokenNV-pIndexTypes-parameter"></a>
  VUID-VkIndirectCommandsLayoutTokenNV-pIndexTypes-parameter  
  If `indexTypeCount` is not `0`, `pIndexTypes` **must** be a valid
  pointer to an array of `indexTypeCount` valid
  [VkIndexType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkIndexType.html) values

- <a
  href="#VUID-VkIndirectCommandsLayoutTokenNV-pIndexTypeValues-parameter"
  id="VUID-VkIndirectCommandsLayoutTokenNV-pIndexTypeValues-parameter"></a>
  VUID-VkIndirectCommandsLayoutTokenNV-pIndexTypeValues-parameter  
  If `indexTypeCount` is not `0`, `pIndexTypeValues` **must** be a valid
  pointer to an array of `indexTypeCount` `uint32_t` values

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_device_generated_commands](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_device_generated_commands.html),
[VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html), [VkIndexType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkIndexType.html),
[VkIndirectCommandsLayoutCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkIndirectCommandsLayoutCreateInfoNV.html),
[VkIndirectCommandsTokenTypeNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkIndirectCommandsTokenTypeNV.html),
[VkIndirectStateFlagsNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkIndirectStateFlagsNV.html),
[VkPipelineLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLayout.html),
[VkShaderStageFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderStageFlags.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkIndirectCommandsLayoutTokenNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
