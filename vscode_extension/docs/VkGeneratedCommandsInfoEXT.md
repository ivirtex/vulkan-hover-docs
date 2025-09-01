# VkGeneratedCommandsInfoEXT(3) Manual Page

## Name

VkGeneratedCommandsInfoEXT - Structure specifying parameters for the generation of commands



## [](#_c_specification)C Specification

The `VkGeneratedCommandsInfoEXT` is defined as:

```c++
// Provided by VK_EXT_device_generated_commands
typedef struct VkGeneratedCommandsInfoEXT {
    VkStructureType                sType;
    const void*                    pNext;
    VkShaderStageFlags             shaderStages;
    VkIndirectExecutionSetEXT      indirectExecutionSet;
    VkIndirectCommandsLayoutEXT    indirectCommandsLayout;
    VkDeviceAddress                indirectAddress;
    VkDeviceSize                   indirectAddressSize;
    VkDeviceAddress                preprocessAddress;
    VkDeviceSize                   preprocessSize;
    uint32_t                       maxSequenceCount;
    VkDeviceAddress                sequenceCountAddress;
    uint32_t                       maxDrawCount;
} VkGeneratedCommandsInfoEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `shaderStages` is the mask of shader stages used by the commands.
- `indirectExecutionSet` is the indirect execution set to be used for binding shaders.
- `indirectCommandsLayout` is the [VkIndirectCommandsLayoutEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndirectCommandsLayoutEXT.html) that specifies the command sequence data.
- `indirectAddress` is an address that holds the indirect buffer data.
- `indirectAddressSize` is the size in bytes of indirect buffer data starting at `indirectAddress`.
- `preprocessAddress` specifies a physical address of the `VkBuffer` used for preprocessing the input data for execution. If this structure is used with [vkCmdExecuteGeneratedCommandsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdExecuteGeneratedCommandsEXT.html) with its `isPreprocessed` set to `VK_TRUE`, then the preprocessing step is skipped but data in this address **may** still be modified. The contents and the layout of this address are opaque to applications and **must** not be modified outside functions related to device-generated commands or copied to another buffer for reuse.
- `preprocessSize` is the maximum byte size within `preprocessAddress` that is available for preprocessing.
- `maxSequenceCount` is used to determine the number of sequences to execute.
- `sequenceCountAddress` specifies an optional physical address of a single `uint32_t` value containing the requested number of sequences to execute.
- `maxDrawCount` is the maximum number of indirect draws that can be executed by any COUNT-type multi-draw indirect tokens. The draw count in the indirect buffer is clamped to this value for these token types.

## [](#_description)Description

If `sequenceCountAddress` is not `NULL`, then `maxSequenceCount` is the maximum number of sequences that can be executed. The actual number is `min(maxSequenceCount, *sequenceCountAddress)`. If `sequenceCountAddress` is `NULL`, then `maxSequenceCount` is the exact number of sequences to execute.

If the action command token for the layout is not a COUNT-type multi-draw indirect token, `maxDrawCount` is ignored.

Valid Usage

- [](#VUID-VkGeneratedCommandsInfoEXT-preprocessAddress-11063)VUID-VkGeneratedCommandsInfoEXT-preprocessAddress-11063  
  If [vkGetGeneratedCommandsMemoryRequirementsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetGeneratedCommandsMemoryRequirementsEXT.html) returns a non-zero size, `preprocessAddress` **must** not be `0`
- [](#VUID-VkGeneratedCommandsInfoEXT-preprocessAddress-11064)VUID-VkGeneratedCommandsInfoEXT-preprocessAddress-11064  
  `VkDeviceMemory` objects bound to the underlying buffer for `preprocessAddress` **must** have been allocated using one of the memory types allowed in the `memoryTypeBits` member of the [VkMemoryRequirements](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryRequirements.html) structure returned by [vkGetGeneratedCommandsMemoryRequirementsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetGeneratedCommandsMemoryRequirementsEXT.html)
- [](#VUID-VkGeneratedCommandsInfoEXT-indirectCommandsLayout-11065)VUID-VkGeneratedCommandsInfoEXT-indirectCommandsLayout-11065  
  If the `indirectCommandsLayout` uses a token of `VK_INDIRECT_COMMANDS_TOKEN_TYPE_PUSH_CONSTANT_EXT`, then the `indirectExecutionSet`’s push constant layout **must** contain the `updateRange` specified in [VkIndirectCommandsPushConstantTokenEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndirectCommandsPushConstantTokenEXT.html)
- [](#VUID-VkGeneratedCommandsInfoEXT-indirectCommandsLayout-11066)VUID-VkGeneratedCommandsInfoEXT-indirectCommandsLayout-11066  
  If the `indirectCommandsLayout` uses a token of `VK_INDIRECT_COMMANDS_TOKEN_TYPE_SEQUENCE_INDEX_EXT`, then the `indirectExecutionSet`’s push constant layout **must** contain the `updateRange` specified in [VkIndirectCommandsPushConstantTokenEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndirectCommandsPushConstantTokenEXT.html)
- [](#VUID-VkGeneratedCommandsInfoEXT-maxSequenceCount-11067)VUID-VkGeneratedCommandsInfoEXT-maxSequenceCount-11067  
  `maxSequenceCount` **must** be less or equal to [VkPhysicalDeviceDeviceGeneratedCommandsPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceDeviceGeneratedCommandsPropertiesEXT.html)::`maxIndirectSequenceCount` and [VkGeneratedCommandsMemoryRequirementsInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGeneratedCommandsMemoryRequirementsInfoEXT.html)::`maxSequenceCount` that was used to determine the `preprocessSize`
- [](#VUID-VkGeneratedCommandsInfoEXT-sequenceCountAddress-11068)VUID-VkGeneratedCommandsInfoEXT-sequenceCountAddress-11068  
  If `sequenceCountAddress` is not `NULL`, the value contained in the address **must** be less or equal to [VkPhysicalDeviceDeviceGeneratedCommandsPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceDeviceGeneratedCommandsPropertiesEXT.html)::`maxIndirectSequenceCount` and [VkGeneratedCommandsMemoryRequirementsInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGeneratedCommandsMemoryRequirementsInfoEXT.html)::`maxSequenceCount` that was used to determine the `preprocessSize`
- [](#VUID-VkGeneratedCommandsInfoEXT-maxSequenceCount-10246)VUID-VkGeneratedCommandsInfoEXT-maxSequenceCount-10246  
  `maxSequenceCount` **must** not be zero
- [](#VUID-VkGeneratedCommandsInfoEXT-preprocessAddress-11069)VUID-VkGeneratedCommandsInfoEXT-preprocessAddress-11069  
  The underlying buffer for `preprocessAddress` **must** have the `VK_BUFFER_USAGE_2_PREPROCESS_BUFFER_BIT_EXT` bit set in its usage flag
- [](#VUID-VkGeneratedCommandsInfoEXT-indirectCommandsLayout-11144)VUID-VkGeneratedCommandsInfoEXT-indirectCommandsLayout-11144  
  If the `indirectCommandsLayout` contains a `VK_INDIRECT_COMMANDS_TOKEN_TYPE_EXECUTION_SET_EXT` token, then the descriptor and push constant layout info provided either by `pipelineLayout` or through a [VkPipelineLayoutCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineLayoutCreateInfo.html) in `pNext` of the [VkIndirectCommandsLayoutCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndirectCommandsLayoutCreateInfoEXT.html) used to create `indirectCommandsLayout` **must** be [compatible](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#descriptorsets-compatibility) with the descriptor and push constant layout info used by `indirectExecutionSet`
- [](#VUID-VkGeneratedCommandsInfoEXT-indirectCommandsLayout-11002)VUID-VkGeneratedCommandsInfoEXT-indirectCommandsLayout-11002  
  If `indirectCommandsLayout` was created with a token sequence that contained the `VK_INDIRECT_COMMANDS_TOKEN_TYPE_EXECUTION_SET_EXT` token, the shader stages used to create the initial shader state of `indirectExecutionSet` **must** equal the [VkIndirectCommandsExecutionSetTokenEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndirectCommandsExecutionSetTokenEXT.html)::`shaderStages` used to create `indirectCommandsLayout`
- [](#VUID-VkGeneratedCommandsInfoEXT-preprocessSize-11071)VUID-VkGeneratedCommandsInfoEXT-preprocessSize-11071  
  `preprocessSize` **must** be greater than or equal to the memory requirement’s size returned by [vkGetGeneratedCommandsMemoryRequirementsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetGeneratedCommandsMemoryRequirementsEXT.html) using the matching inputs (`indirectCommandsLayout`, …​) as within this structure
- [](#VUID-VkGeneratedCommandsInfoEXT-sequenceCountAddress-11072)VUID-VkGeneratedCommandsInfoEXT-sequenceCountAddress-11072  
  The underlying buffer for `sequenceCountAddress` **must** have the `VK_BUFFER_USAGE_2_INDIRECT_BUFFER_BIT_KHR` bit set in its usage flag
- [](#VUID-VkGeneratedCommandsInfoEXT-sequenceCountAddress-11073)VUID-VkGeneratedCommandsInfoEXT-sequenceCountAddress-11073  
  If `sequenceCountAddress` is not `NULL`, `sequenceCountAddress` **must** be aligned to `4`
- [](#VUID-VkGeneratedCommandsInfoEXT-indirectAddress-11074)VUID-VkGeneratedCommandsInfoEXT-indirectAddress-11074  
  `indirectAddress` **must** be aligned to `4`
- [](#VUID-VkGeneratedCommandsInfoEXT-indirectAddressSize-11077)VUID-VkGeneratedCommandsInfoEXT-indirectAddressSize-11077  
  `indirectAddressSize` **must** be greater than zero
- [](#VUID-VkGeneratedCommandsInfoEXT-maxDrawCount-11078)VUID-VkGeneratedCommandsInfoEXT-maxDrawCount-11078  
  When not ignored, `maxDrawCount` × `maxSequenceCount` **must** be less than 2^24
- [](#VUID-VkGeneratedCommandsInfoEXT-indirectCommandsLayout-11079)VUID-VkGeneratedCommandsInfoEXT-indirectCommandsLayout-11079  
  If `indirectCommandsLayout` was created using a `VK_INDIRECT_COMMANDS_TOKEN_TYPE_VERTEX_BUFFER_EXT` token and shader objects are not bound then the bound graphics pipeline **must** have been created with `VK_DYNAMIC_STATE_VERTEX_INPUT_BINDING_STRIDE` in `pDynamicStates`
- [](#VUID-VkGeneratedCommandsInfoEXT-indirectCommandsLayout-11083)VUID-VkGeneratedCommandsInfoEXT-indirectCommandsLayout-11083  
  If the token sequence of the passed `indirectCommandsLayout` contains a `VK_INDIRECT_COMMANDS_TOKEN_TYPE_EXECUTION_SET_EXT` token, the `indirectExecutionSet` **must** not be [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html)
- [](#VUID-VkGeneratedCommandsInfoEXT-indirectCommandsLayout-10241)VUID-VkGeneratedCommandsInfoEXT-indirectCommandsLayout-10241  
  If the token sequence of the passed `indirectCommandsLayout` does not contains a `VK_INDIRECT_COMMANDS_TOKEN_TYPE_EXECUTION_SET_EXT` token, the `indirectExecutionSet` **must** be [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html)
- [](#VUID-VkGeneratedCommandsInfoEXT-indirectExecutionSet-11080)VUID-VkGeneratedCommandsInfoEXT-indirectExecutionSet-11080  
  If `indirectExecutionSet` is [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), a [VkGeneratedCommandsPipelineInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGeneratedCommandsPipelineInfoEXT.html) or [VkGeneratedCommandsShaderInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGeneratedCommandsShaderInfoEXT.html) **must** be included in the `pNext` chain

Valid Usage (Implicit)

- [](#VUID-VkGeneratedCommandsInfoEXT-sType-sType)VUID-VkGeneratedCommandsInfoEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_GENERATED_COMMANDS_INFO_EXT`
- [](#VUID-VkGeneratedCommandsInfoEXT-shaderStages-parameter)VUID-VkGeneratedCommandsInfoEXT-shaderStages-parameter  
  `shaderStages` **must** be a valid combination of [VkShaderStageFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderStageFlagBits.html) values
- [](#VUID-VkGeneratedCommandsInfoEXT-shaderStages-requiredbitmask)VUID-VkGeneratedCommandsInfoEXT-shaderStages-requiredbitmask  
  `shaderStages` **must** not be `0`
- [](#VUID-VkGeneratedCommandsInfoEXT-indirectExecutionSet-parameter)VUID-VkGeneratedCommandsInfoEXT-indirectExecutionSet-parameter  
  If `indirectExecutionSet` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), `indirectExecutionSet` **must** be a valid [VkIndirectExecutionSetEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndirectExecutionSetEXT.html) handle
- [](#VUID-VkGeneratedCommandsInfoEXT-indirectCommandsLayout-parameter)VUID-VkGeneratedCommandsInfoEXT-indirectCommandsLayout-parameter  
  `indirectCommandsLayout` **must** be a valid [VkIndirectCommandsLayoutEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndirectCommandsLayoutEXT.html) handle
- [](#VUID-VkGeneratedCommandsInfoEXT-indirectAddress-parameter)VUID-VkGeneratedCommandsInfoEXT-indirectAddress-parameter  
  `indirectAddress` **must** be a valid [VkDeviceAddress](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceAddress.html) value
- [](#VUID-VkGeneratedCommandsInfoEXT-preprocessAddress-parameter)VUID-VkGeneratedCommandsInfoEXT-preprocessAddress-parameter  
  If `preprocessAddress` is not `0`, `preprocessAddress` **must** be a valid [VkDeviceAddress](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceAddress.html) value
- [](#VUID-VkGeneratedCommandsInfoEXT-sequenceCountAddress-parameter)VUID-VkGeneratedCommandsInfoEXT-sequenceCountAddress-parameter  
  If `sequenceCountAddress` is not `0`, `sequenceCountAddress` **must** be a valid [VkDeviceAddress](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceAddress.html) value
- [](#VUID-VkGeneratedCommandsInfoEXT-commonparent)VUID-VkGeneratedCommandsInfoEXT-commonparent  
  Both of `indirectCommandsLayout`, and `indirectExecutionSet` that are valid handles of non-ignored parameters **must** have been created, allocated, or retrieved from the same [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html)

## [](#_see_also)See Also

[VK\_EXT\_device\_generated\_commands](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_device_generated_commands.html), [VkDeviceAddress](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceAddress.html), [VkDeviceSize](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceSize.html), [VkIndirectCommandsLayoutEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndirectCommandsLayoutEXT.html), [VkIndirectExecutionSetEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndirectExecutionSetEXT.html), [VkShaderStageFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderStageFlags.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkCmdExecuteGeneratedCommandsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdExecuteGeneratedCommandsEXT.html), [vkCmdPreprocessGeneratedCommandsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdPreprocessGeneratedCommandsEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkGeneratedCommandsInfoEXT).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0