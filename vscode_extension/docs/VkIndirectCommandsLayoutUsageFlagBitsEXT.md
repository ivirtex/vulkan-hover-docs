# VkIndirectCommandsLayoutUsageFlagBitsEXT(3) Manual Page

## Name

VkIndirectCommandsLayoutUsageFlagBitsEXT - Bitmask specifying allowed usage of an indirect commands layout



## [](#_c_specification)C Specification

Bits which **can** be set in [VkIndirectCommandsLayoutCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndirectCommandsLayoutCreateInfoEXT.html)::`flags`, specifying usage hints of an indirect command layout, are:

```c++
// Provided by VK_EXT_device_generated_commands
typedef enum VkIndirectCommandsLayoutUsageFlagBitsEXT {
    VK_INDIRECT_COMMANDS_LAYOUT_USAGE_EXPLICIT_PREPROCESS_BIT_EXT = 0x00000001,
    VK_INDIRECT_COMMANDS_LAYOUT_USAGE_UNORDERED_SEQUENCES_BIT_EXT = 0x00000002,
} VkIndirectCommandsLayoutUsageFlagBitsEXT;
```

## [](#_description)Description

- `VK_INDIRECT_COMMANDS_LAYOUT_USAGE_EXPLICIT_PREPROCESS_BIT_EXT` specifies that the layout is always used with the manual preprocessing step through calling [vkCmdPreprocessGeneratedCommandsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdPreprocessGeneratedCommandsEXT.html) and executed by [vkCmdExecuteGeneratedCommandsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdExecuteGeneratedCommandsEXT.html) with `isPreprocessed` set to `VK_TRUE`.
- `VK_INDIRECT_COMMANDS_LAYOUT_USAGE_UNORDERED_SEQUENCES_BIT_EXT` specifies that the processing of sequences will happen at an implementation-dependent order, which is not guaranteed to be deterministic using the same input data. This flag is ignored when the `shaderStages` is `VK_SHADER_STAGE_COMPUTE_BIT` as it is implied that the dispatch sequence is always unordered.

## [](#_see_also)See Also

[VK\_EXT\_device\_generated\_commands](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_device_generated_commands.html), [VkIndirectCommandsLayoutUsageFlagsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndirectCommandsLayoutUsageFlagsEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkIndirectCommandsLayoutUsageFlagBitsEXT).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0