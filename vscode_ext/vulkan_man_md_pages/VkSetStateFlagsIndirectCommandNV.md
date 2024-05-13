# VkSetStateFlagsIndirectCommandNV(3) Manual Page

## Name

VkSetStateFlagsIndirectCommandNV - Structure specifying input data for a
single state flag command token



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkSetStateFlagsIndirectCommandNV` structure specifies the input
data for the `VK_INDIRECT_COMMANDS_TOKEN_TYPE_STATE_FLAGS_NV` token.
Which state is changed depends on the
[VkIndirectStateFlagBitsNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkIndirectStateFlagBitsNV.html) specified at
`VkIndirectCommandsLayoutNV` creation time.

``` c
// Provided by VK_NV_device_generated_commands
typedef struct VkSetStateFlagsIndirectCommandNV {
    uint32_t    data;
} VkSetStateFlagsIndirectCommandNV;
```

## <a href="#_members" class="anchor"></a>Members

- `data` encodes packed state that this command alters.

  - Bit `0`: If set represents `VK_FRONT_FACE_CLOCKWISE`, otherwise
    `VK_FRONT_FACE_COUNTER_CLOCKWISE`

## <a href="#_description" class="anchor"></a>Description

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_device_generated_commands](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_device_generated_commands.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkSetStateFlagsIndirectCommandNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
