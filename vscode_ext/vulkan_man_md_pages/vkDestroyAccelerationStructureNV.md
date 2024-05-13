# vkDestroyAccelerationStructureNV(3) Manual Page

## Name

vkDestroyAccelerationStructureNV - Destroy an acceleration structure
object



## <a href="#_c_specification" class="anchor"></a>C Specification

To destroy an acceleration structure, call:

``` c
// Provided by VK_NV_ray_tracing
void vkDestroyAccelerationStructureNV(
    VkDevice                                    device,
    VkAccelerationStructureNV                   accelerationStructure,
    const VkAllocationCallbacks*                pAllocator);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device that destroys the buffer.

- `accelerationStructure` is the acceleration structure to destroy.

- `pAllocator` controls host memory allocation as described in the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#memory-allocation"
  target="_blank" rel="noopener">Memory Allocation</a> chapter.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a
  href="#VUID-vkDestroyAccelerationStructureNV-accelerationStructure-03752"
  id="VUID-vkDestroyAccelerationStructureNV-accelerationStructure-03752"></a>
  VUID-vkDestroyAccelerationStructureNV-accelerationStructure-03752  
  All submitted commands that refer to `accelerationStructure` **must**
  have completed execution

- <a
  href="#VUID-vkDestroyAccelerationStructureNV-accelerationStructure-03753"
  id="VUID-vkDestroyAccelerationStructureNV-accelerationStructure-03753"></a>
  VUID-vkDestroyAccelerationStructureNV-accelerationStructure-03753  
  If `VkAllocationCallbacks` were provided when `accelerationStructure`
  was created, a compatible set of callbacks **must** be provided here

- <a
  href="#VUID-vkDestroyAccelerationStructureNV-accelerationStructure-03754"
  id="VUID-vkDestroyAccelerationStructureNV-accelerationStructure-03754"></a>
  VUID-vkDestroyAccelerationStructureNV-accelerationStructure-03754  
  If no `VkAllocationCallbacks` were provided when
  `accelerationStructure` was created, `pAllocator` **must** be `NULL`

Valid Usage (Implicit)

- <a href="#VUID-vkDestroyAccelerationStructureNV-device-parameter"
  id="VUID-vkDestroyAccelerationStructureNV-device-parameter"></a>
  VUID-vkDestroyAccelerationStructureNV-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a
  href="#VUID-vkDestroyAccelerationStructureNV-accelerationStructure-parameter"
  id="VUID-vkDestroyAccelerationStructureNV-accelerationStructure-parameter"></a>
  VUID-vkDestroyAccelerationStructureNV-accelerationStructure-parameter  
  If `accelerationStructure` is not
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), `accelerationStructure`
  **must** be a valid
  [VkAccelerationStructureNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureNV.html) handle

- <a href="#VUID-vkDestroyAccelerationStructureNV-pAllocator-parameter"
  id="VUID-vkDestroyAccelerationStructureNV-pAllocator-parameter"></a>
  VUID-vkDestroyAccelerationStructureNV-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid
  pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html)
  structure

- <a
  href="#VUID-vkDestroyAccelerationStructureNV-accelerationStructure-parent"
  id="VUID-vkDestroyAccelerationStructureNV-accelerationStructure-parent"></a>
  VUID-vkDestroyAccelerationStructureNV-accelerationStructure-parent  
  If `accelerationStructure` is a valid handle, it **must** have been
  created, allocated, or retrieved from `device`

Host Synchronization

- Host access to `accelerationStructure` **must** be externally
  synchronized

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_ray_tracing](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_ray_tracing.html),
[VkAccelerationStructureNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureNV.html),
[VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkDestroyAccelerationStructureNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
