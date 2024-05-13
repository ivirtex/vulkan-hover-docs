# VkIndirectCommandsLayoutNV(3) Manual Page

## Name

VkIndirectCommandsLayoutNV - Opaque handle to an indirect commands
layout object



## <a href="#_c_specification" class="anchor"></a>C Specification

The device-side command generation happens through an iterative
processing of an atomic sequence comprised of command tokens, which are
represented by:

``` c
// Provided by VK_NV_device_generated_commands
VK_DEFINE_NON_DISPATCHABLE_HANDLE(VkIndirectCommandsLayoutNV)
```

## <a href="#_description" class="anchor"></a>Description

Each indirect command layout **must** have exactly one action command
token and it **must** be the last token in the sequence.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_device_generated_commands](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_device_generated_commands.html),
[VkGeneratedCommandsInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGeneratedCommandsInfoNV.html),
[VkGeneratedCommandsMemoryRequirementsInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGeneratedCommandsMemoryRequirementsInfoNV.html),
[vkCreateIndirectCommandsLayoutNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateIndirectCommandsLayoutNV.html),
[vkDestroyIndirectCommandsLayoutNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkDestroyIndirectCommandsLayoutNV.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkIndirectCommandsLayoutNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
