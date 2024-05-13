# VkDispatchIndirectCommand(3) Manual Page

## Name

VkDispatchIndirectCommand - Structure specifying a indirect dispatching
command



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkDispatchIndirectCommand` structure is defined as:

``` c
// Provided by VK_VERSION_1_0
typedef struct VkDispatchIndirectCommand {
    uint32_t    x;
    uint32_t    y;
    uint32_t    z;
} VkDispatchIndirectCommand;
```

## <a href="#_members" class="anchor"></a>Members

- `x` is the number of local workgroups to dispatch in the X dimension.

- `y` is the number of local workgroups to dispatch in the Y dimension.

- `z` is the number of local workgroups to dispatch in the Z dimension.

## <a href="#_description" class="anchor"></a>Description

The members of `VkDispatchIndirectCommand` have the same meaning as the
corresponding parameters of [vkCmdDispatch](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdDispatch.html).

Valid Usage

- <a href="#VUID-VkDispatchIndirectCommand-x-00417"
  id="VUID-VkDispatchIndirectCommand-x-00417"></a>
  VUID-VkDispatchIndirectCommand-x-00417  
  `x` **must** be less than or equal to
  `VkPhysicalDeviceLimits`::`maxComputeWorkGroupCount`\[0\]

- <a href="#VUID-VkDispatchIndirectCommand-y-00418"
  id="VUID-VkDispatchIndirectCommand-y-00418"></a>
  VUID-VkDispatchIndirectCommand-y-00418  
  `y` **must** be less than or equal to
  `VkPhysicalDeviceLimits`::`maxComputeWorkGroupCount`\[1\]

- <a href="#VUID-VkDispatchIndirectCommand-z-00419"
  id="VUID-VkDispatchIndirectCommand-z-00419"></a>
  VUID-VkDispatchIndirectCommand-z-00419  
  `z` **must** be less than or equal to
  `VkPhysicalDeviceLimits`::`maxComputeWorkGroupCount`\[2\]

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[vkCmdDispatchIndirect](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdDispatchIndirect.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkDispatchIndirectCommand"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
