# VkDispatchGraphCountInfoAMDX(3) Manual Page

## Name

VkDispatchGraphCountInfoAMDX - Structure specifying count parameters for execution graph dispatch



## [](#_c_specification)C Specification

The `VkDispatchGraphCountInfoAMDX` structure is defined as:

```c++
// Provided by VK_AMDX_shader_enqueue
typedef struct VkDispatchGraphCountInfoAMDX {
    uint32_t                          count;
    VkDeviceOrHostAddressConstAMDX    infos;
    uint64_t                          stride;
} VkDispatchGraphCountInfoAMDX;
```

## [](#_members)Members

- `count` is the number of dispatches to perform.
- `infos` is the device or host address of a flat array of [VkDispatchGraphInfoAMDX](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDispatchGraphInfoAMDX.html) structures
- `stride` is the byte stride between successive [VkDispatchGraphInfoAMDX](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDispatchGraphInfoAMDX.html) structures in `infos`

## [](#_description)Description

Whether `infos` is consumed as a device or host pointer is defined by the command this structure is used in.

## [](#_see_also)See Also

[VK\_AMDX\_shader\_enqueue](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_AMDX_shader_enqueue.html), [VkDeviceOrHostAddressConstAMDX](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceOrHostAddressConstAMDX.html), [vkCmdDispatchGraphAMDX](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdDispatchGraphAMDX.html), [vkCmdDispatchGraphIndirectAMDX](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdDispatchGraphIndirectAMDX.html), [vkCmdDispatchGraphIndirectCountAMDX](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdDispatchGraphIndirectCountAMDX.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkDispatchGraphCountInfoAMDX)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0