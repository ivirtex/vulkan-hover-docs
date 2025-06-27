# VkIndirectCommandsLayoutUsageFlagBitsNV(3) Manual Page

## Name

VkIndirectCommandsLayoutUsageFlagBitsNV - Bitmask specifying allowed
usage of an indirect commands layout



## <a href="#_c_specification" class="anchor"></a>C Specification

Bits which **can** be set in
[VkIndirectCommandsLayoutCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkIndirectCommandsLayoutCreateInfoNV.html)::`flags`,
specifying usage hints of an indirect command layout, are:

``` c
// Provided by VK_NV_device_generated_commands
typedef enum VkIndirectCommandsLayoutUsageFlagBitsNV {
    VK_INDIRECT_COMMANDS_LAYOUT_USAGE_EXPLICIT_PREPROCESS_BIT_NV = 0x00000001,
    VK_INDIRECT_COMMANDS_LAYOUT_USAGE_INDEXED_SEQUENCES_BIT_NV = 0x00000002,
    VK_INDIRECT_COMMANDS_LAYOUT_USAGE_UNORDERED_SEQUENCES_BIT_NV = 0x00000004,
} VkIndirectCommandsLayoutUsageFlagBitsNV;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_INDIRECT_COMMANDS_LAYOUT_USAGE_EXPLICIT_PREPROCESS_BIT_NV`
  specifies that the layout is always used with the manual preprocessing
  step through calling
  [vkCmdPreprocessGeneratedCommandsNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdPreprocessGeneratedCommandsNV.html)
  and executed by
  [vkCmdExecuteGeneratedCommandsNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdExecuteGeneratedCommandsNV.html)
  with `isPreprocessed` set to `VK_TRUE`.

- `VK_INDIRECT_COMMANDS_LAYOUT_USAGE_INDEXED_SEQUENCES_BIT_NV` specifies
  that the input data for the sequences is not implicitly indexed from
  0..sequencesUsed, but an application-provided `VkBuffer` encoding the
  index is provided.

- `VK_INDIRECT_COMMANDS_LAYOUT_USAGE_UNORDERED_SEQUENCES_BIT_NV`
  specifies that the processing of sequences **can** happen at an
  implementation-dependent order, which is not: guaranteed to be
  coherent using the same input data. This flag is ignored when the
  `pipelineBindPoint` is `VK_PIPELINE_BIND_POINT_COMPUTE` as it is
  implied that the dispatch sequence is always unordered.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_device_generated_commands](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_device_generated_commands.html),
[VkIndirectCommandsLayoutUsageFlagsNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkIndirectCommandsLayoutUsageFlagsNV.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkIndirectCommandsLayoutUsageFlagBitsNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
