# VkSetStateFlagsIndirectCommandNV(3) Manual Page

## Name

VkSetStateFlagsIndirectCommandNV - Structure specifying input data for a single state flag command token



## [](#_c_specification)C Specification

The `VkSetStateFlagsIndirectCommandNV` structure specifies the input data for the `VK_INDIRECT_COMMANDS_TOKEN_TYPE_STATE_FLAGS_NV` token. Which state is changed depends on the [VkIndirectStateFlagBitsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndirectStateFlagBitsNV.html) specified at `VkIndirectCommandsLayoutNV` creation time.

```c++
// Provided by VK_NV_device_generated_commands
typedef struct VkSetStateFlagsIndirectCommandNV {
    uint32_t    data;
} VkSetStateFlagsIndirectCommandNV;
```

## [](#_members)Members

- `data` encodes packed state that this command alters.
  
  - Bit `0`: If set represents `VK_FRONT_FACE_CLOCKWISE`, otherwise `VK_FRONT_FACE_COUNTER_CLOCKWISE`

## [](#_description)Description

## [](#_see_also)See Also

[VK\_NV\_device\_generated\_commands](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_device_generated_commands.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkSetStateFlagsIndirectCommandNV)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0