# VkCommandPoolResetFlagBits(3) Manual Page

## Name

VkCommandPoolResetFlagBits - Bitmask controlling behavior of a command
pool reset



## <a href="#_c_specification" class="anchor"></a>C Specification

Bits which **can** be set in
[vkResetCommandPool](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkResetCommandPool.html)::`flags`, controlling the
reset operation, are:

``` c
// Provided by VK_VERSION_1_0
typedef enum VkCommandPoolResetFlagBits {
    VK_COMMAND_POOL_RESET_RELEASE_RESOURCES_BIT = 0x00000001,
} VkCommandPoolResetFlagBits;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_COMMAND_POOL_RESET_RELEASE_RESOURCES_BIT` specifies that resetting
  a command pool recycles all of the resources from the command pool
  back to the system.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkCommandPoolResetFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandPoolResetFlags.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkCommandPoolResetFlagBits"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
