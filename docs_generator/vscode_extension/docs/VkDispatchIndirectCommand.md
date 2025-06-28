# VkDispatchIndirectCommand(3) Manual Page

## Name

VkDispatchIndirectCommand - Structure specifying an indirect dispatching command



## [](#_c_specification)C Specification

The `VkDispatchIndirectCommand` structure is defined as:

```c++
// Provided by VK_VERSION_1_0
typedef struct VkDispatchIndirectCommand {
    uint32_t    x;
    uint32_t    y;
    uint32_t    z;
} VkDispatchIndirectCommand;
```

## [](#_members)Members

- `x` is the number of local workgroups to dispatch in the X dimension.
- `y` is the number of local workgroups to dispatch in the Y dimension.
- `z` is the number of local workgroups to dispatch in the Z dimension.

## [](#_description)Description

The members of `VkDispatchIndirectCommand` have the same meaning as the corresponding parameters of [vkCmdDispatch](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdDispatch.html).

Valid Usage

- [](#VUID-VkDispatchIndirectCommand-x-00417)VUID-VkDispatchIndirectCommand-x-00417  
  `x` **must** be less than or equal to `VkPhysicalDeviceLimits`::`maxComputeWorkGroupCount`\[0]
- [](#VUID-VkDispatchIndirectCommand-y-00418)VUID-VkDispatchIndirectCommand-y-00418  
  `y` **must** be less than or equal to `VkPhysicalDeviceLimits`::`maxComputeWorkGroupCount`\[1]
- [](#VUID-VkDispatchIndirectCommand-z-00419)VUID-VkDispatchIndirectCommand-z-00419  
  `z` **must** be less than or equal to `VkPhysicalDeviceLimits`::`maxComputeWorkGroupCount`\[2]

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [vkCmdDispatchIndirect](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdDispatchIndirect.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkDispatchIndirectCommand)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0