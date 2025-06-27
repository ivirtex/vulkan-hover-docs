# vkDestroyInstance(3) Manual Page

## Name

vkDestroyInstance - Destroy an instance of Vulkan



## <a href="#_c_specification" class="anchor"></a>C Specification

To destroy an instance, call:

``` c
// Provided by VK_VERSION_1_0
void vkDestroyInstance(
    VkInstance                                  instance,
    const VkAllocationCallbacks*                pAllocator);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `instance` is the handle of the instance to destroy.

- `pAllocator` controls host memory allocation as described in the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#memory-allocation"
  target="_blank" rel="noopener">Memory Allocation</a> chapter.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-vkDestroyInstance-instance-00629"
  id="VUID-vkDestroyInstance-instance-00629"></a>
  VUID-vkDestroyInstance-instance-00629  
  All child objects created using `instance` **must** have been
  destroyed prior to destroying `instance`

- <a href="#VUID-vkDestroyInstance-instance-00630"
  id="VUID-vkDestroyInstance-instance-00630"></a>
  VUID-vkDestroyInstance-instance-00630  
  If `VkAllocationCallbacks` were provided when `instance` was created,
  a compatible set of callbacks **must** be provided here

- <a href="#VUID-vkDestroyInstance-instance-00631"
  id="VUID-vkDestroyInstance-instance-00631"></a>
  VUID-vkDestroyInstance-instance-00631  
  If no `VkAllocationCallbacks` were provided when `instance` was
  created, `pAllocator` **must** be `NULL`

Valid Usage (Implicit)

- <a href="#VUID-vkDestroyInstance-instance-parameter"
  id="VUID-vkDestroyInstance-instance-parameter"></a>
  VUID-vkDestroyInstance-instance-parameter  
  If `instance` is not `NULL`, `instance` **must** be a valid
  [VkInstance](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkInstance.html) handle

- <a href="#VUID-vkDestroyInstance-pAllocator-parameter"
  id="VUID-vkDestroyInstance-pAllocator-parameter"></a>
  VUID-vkDestroyInstance-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid
  pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html)
  structure

Host Synchronization

- Host access to `instance` **must** be externally synchronized

- Host access to all `VkPhysicalDevice` objects enumerated from
  `instance` **must** be externally synchronized

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html),
[VkInstance](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkInstance.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkDestroyInstance"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
