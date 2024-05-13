# VkPhysicalDeviceNestedCommandBufferFeaturesEXT(3) Manual Page

## Name

VkPhysicalDeviceNestedCommandBufferFeaturesEXT - Structure describing
whether nested command buffers are supported by the implementation



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceNestedCommandBufferFeaturesEXT` structure is
defined as:

``` c
// Provided by VK_EXT_nested_command_buffer
typedef struct VkPhysicalDeviceNestedCommandBufferFeaturesEXT {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           nestedCommandBuffer;
    VkBool32           nestedCommandBufferRendering;
    VkBool32           nestedCommandBufferSimultaneousUse;
} VkPhysicalDeviceNestedCommandBufferFeaturesEXT;
```

## <a href="#_members" class="anchor"></a>Members

This structure describes the following features:

- <span id="features-nestedCommandBuffer"></span> `nestedCommandBuffer`
  indicates the implementation supports nested command buffers, which
  allows <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#glossary"
  target="_blank" rel="noopener">Secondary Command Buffers</a> to
  execute other <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#glossary"
  target="_blank" rel="noopener">Secondary Command Buffers</a>.

- <span id="features-nestedCommandBufferRendering"></span>
  `nestedCommandBufferRendering` indicates that it is valid to call
  [vkCmdExecuteCommands](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdExecuteCommands.html) inside a <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#glossary"
  target="_blank" rel="noopener">Secondary Command Buffer</a> recorded
  with `VK_COMMAND_BUFFER_USAGE_RENDER_PASS_CONTINUE_BIT`.

- <span id="features-nestedCommandBufferSimultaneousUse"></span>
  `nestedCommandBufferSimultaneousUse` indicates that the implementation
  supports nested command buffers with command buffers that are recorded
  with `VK_COMMAND_BUFFER_USAGE_SIMULTANEOUS_USE_BIT`.

## <a href="#_description" class="anchor"></a>Description

If the `VkPhysicalDeviceNestedCommandBufferFeaturesEXT` structure is
included in the `pNext` chain of the
[VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html) structure
passed to
[vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceFeatures2.html), it is
filled in to indicate whether each corresponding feature is supported.
`VkPhysicalDeviceNestedCommandBufferFeaturesEXT` **can** also be used in
the `pNext` chain of [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html) to
selectively enable these features.

Valid Usage (Implicit)

- <a
  href="#VUID-VkPhysicalDeviceNestedCommandBufferFeaturesEXT-sType-sType"
  id="VUID-VkPhysicalDeviceNestedCommandBufferFeaturesEXT-sType-sType"></a>
  VUID-VkPhysicalDeviceNestedCommandBufferFeaturesEXT-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_NESTED_COMMAND_BUFFER_FEATURES_EXT`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_nested_command_buffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_nested_command_buffer.html),
[VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceNestedCommandBufferFeaturesEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
