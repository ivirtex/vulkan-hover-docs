# vkDestroyIndirectCommandsLayoutNV(3) Manual Page

## Name

vkDestroyIndirectCommandsLayoutNV - Destroy an indirect commands layout



## <a href="#_c_specification" class="anchor"></a>C Specification

Indirect command layouts are destroyed by:

``` c
// Provided by VK_NV_device_generated_commands
void vkDestroyIndirectCommandsLayoutNV(
    VkDevice                                    device,
    VkIndirectCommandsLayoutNV                  indirectCommandsLayout,
    const VkAllocationCallbacks*                pAllocator);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device that destroys the layout.

- `indirectCommandsLayout` is the layout to destroy.

- `pAllocator` controls host memory allocation as described in the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#memory-allocation"
  target="_blank" rel="noopener">Memory Allocation</a> chapter.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a
  href="#VUID-vkDestroyIndirectCommandsLayoutNV-indirectCommandsLayout-02938"
  id="VUID-vkDestroyIndirectCommandsLayoutNV-indirectCommandsLayout-02938"></a>
  VUID-vkDestroyIndirectCommandsLayoutNV-indirectCommandsLayout-02938  
  All submitted commands that refer to `indirectCommandsLayout` **must**
  have completed execution

- <a
  href="#VUID-vkDestroyIndirectCommandsLayoutNV-indirectCommandsLayout-02939"
  id="VUID-vkDestroyIndirectCommandsLayoutNV-indirectCommandsLayout-02939"></a>
  VUID-vkDestroyIndirectCommandsLayoutNV-indirectCommandsLayout-02939  
  If `VkAllocationCallbacks` were provided when `indirectCommandsLayout`
  was created, a compatible set of callbacks **must** be provided here

- <a
  href="#VUID-vkDestroyIndirectCommandsLayoutNV-indirectCommandsLayout-02940"
  id="VUID-vkDestroyIndirectCommandsLayoutNV-indirectCommandsLayout-02940"></a>
  VUID-vkDestroyIndirectCommandsLayoutNV-indirectCommandsLayout-02940  
  If no `VkAllocationCallbacks` were provided when
  `indirectCommandsLayout` was created, `pAllocator` **must** be `NULL`

- <a
  href="#VUID-vkDestroyIndirectCommandsLayoutNV-deviceGeneratedCommands-02941"
  id="VUID-vkDestroyIndirectCommandsLayoutNV-deviceGeneratedCommands-02941"></a>
  VUID-vkDestroyIndirectCommandsLayoutNV-deviceGeneratedCommands-02941  
  The <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-deviceGeneratedCommands"
  target="_blank"
  rel="noopener"><code>VkPhysicalDeviceDeviceGeneratedCommandsFeaturesNV</code>::<code>deviceGeneratedCommands</code></a>
  feature **must** be enabled

Valid Usage (Implicit)

- <a href="#VUID-vkDestroyIndirectCommandsLayoutNV-device-parameter"
  id="VUID-vkDestroyIndirectCommandsLayoutNV-device-parameter"></a>
  VUID-vkDestroyIndirectCommandsLayoutNV-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a
  href="#VUID-vkDestroyIndirectCommandsLayoutNV-indirectCommandsLayout-parameter"
  id="VUID-vkDestroyIndirectCommandsLayoutNV-indirectCommandsLayout-parameter"></a>
  VUID-vkDestroyIndirectCommandsLayoutNV-indirectCommandsLayout-parameter  
  If `indirectCommandsLayout` is not
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), `indirectCommandsLayout`
  **must** be a valid
  [VkIndirectCommandsLayoutNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkIndirectCommandsLayoutNV.html) handle

- <a href="#VUID-vkDestroyIndirectCommandsLayoutNV-pAllocator-parameter"
  id="VUID-vkDestroyIndirectCommandsLayoutNV-pAllocator-parameter"></a>
  VUID-vkDestroyIndirectCommandsLayoutNV-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid
  pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html)
  structure

- <a
  href="#VUID-vkDestroyIndirectCommandsLayoutNV-indirectCommandsLayout-parent"
  id="VUID-vkDestroyIndirectCommandsLayoutNV-indirectCommandsLayout-parent"></a>
  VUID-vkDestroyIndirectCommandsLayoutNV-indirectCommandsLayout-parent  
  If `indirectCommandsLayout` is a valid handle, it **must** have been
  created, allocated, or retrieved from `device`

Host Synchronization

- Host access to `indirectCommandsLayout` **must** be externally
  synchronized

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_device_generated_commands](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_device_generated_commands.html),
[VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html),
[VkIndirectCommandsLayoutNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkIndirectCommandsLayoutNV.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkDestroyIndirectCommandsLayoutNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
