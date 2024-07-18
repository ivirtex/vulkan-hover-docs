# VkIndirectStateFlagBitsNV(3) Manual Page

## Name

VkIndirectStateFlagBitsNV - Bitmask specifying state that can be altered
on the device



## <a href="#_c_specification" class="anchor"></a>C Specification

A subset of the graphics pipeline state **can** be altered using
indirect state flags:

``` c
// Provided by VK_NV_device_generated_commands
typedef enum VkIndirectStateFlagBitsNV {
    VK_INDIRECT_STATE_FLAG_FRONTFACE_BIT_NV = 0x00000001,
} VkIndirectStateFlagBitsNV;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_INDIRECT_STATE_FLAG_FRONTFACE_BIT_NV` allows to toggle the
  [VkFrontFace](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFrontFace.html) rasterization state for subsequent
  drawing commands.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_device_generated_commands](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_device_generated_commands.html),
[VkIndirectStateFlagsNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkIndirectStateFlagsNV.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkIndirectStateFlagBitsNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
