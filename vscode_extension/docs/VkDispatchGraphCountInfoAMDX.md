# VkDispatchGraphCountInfoAMDX(3) Manual Page

## Name

VkDispatchGraphCountInfoAMDX - Structure specifying count parameters for
execution graph dispatch



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkDispatchGraphCountInfoAMDX` structure is defined as:

``` c
// Provided by VK_AMDX_shader_enqueue
typedef struct VkDispatchGraphCountInfoAMDX {
    uint32_t                          count;
    VkDeviceOrHostAddressConstAMDX    infos;
    uint64_t                          stride;
} VkDispatchGraphCountInfoAMDX;
```

## <a href="#_members" class="anchor"></a>Members

- `count` is the number of dispatches to perform.

- `infos` is the device or host address of a flat array of
  [VkDispatchGraphInfoAMDX](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDispatchGraphInfoAMDX.html) structures

- `stride` is the byte stride between successive
  [VkDispatchGraphInfoAMDX](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDispatchGraphInfoAMDX.html) structures in
  `infos`

## <a href="#_description" class="anchor"></a>Description

Whether `infos` is consumed as a device or host pointer is defined by
the command this structure is used in.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_AMDX_shader_enqueue](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_AMDX_shader_enqueue.html),
[VkDeviceOrHostAddressConstAMDX](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceOrHostAddressConstAMDX.html),
[vkCmdDispatchGraphAMDX](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdDispatchGraphAMDX.html),
[vkCmdDispatchGraphIndirectAMDX](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdDispatchGraphIndirectAMDX.html),
[vkCmdDispatchGraphIndirectCountAMDX](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdDispatchGraphIndirectCountAMDX.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkDispatchGraphCountInfoAMDX"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
