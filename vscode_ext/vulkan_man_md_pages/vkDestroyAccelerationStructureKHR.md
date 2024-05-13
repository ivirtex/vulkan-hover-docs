# vkDestroyAccelerationStructureKHR(3) Manual Page

## Name

vkDestroyAccelerationStructureKHR - Destroy an acceleration structure
object



## <a href="#_c_specification" class="anchor"></a>C Specification

To destroy an acceleration structure, call:

``` c
// Provided by VK_KHR_acceleration_structure
void vkDestroyAccelerationStructureKHR(
    VkDevice                                    device,
    VkAccelerationStructureKHR                  accelerationStructure,
    const VkAllocationCallbacks*                pAllocator);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device that destroys the acceleration
  structure.

- `accelerationStructure` is the acceleration structure to destroy.

- `pAllocator` controls host memory allocation as described in the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#memory-allocation"
  target="_blank" rel="noopener">Memory Allocation</a> chapter.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a
  href="#VUID-vkDestroyAccelerationStructureKHR-accelerationStructure-08934"
  id="VUID-vkDestroyAccelerationStructureKHR-accelerationStructure-08934"></a>
  VUID-vkDestroyAccelerationStructureKHR-accelerationStructure-08934  
  The <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-accelerationStructure"
  target="_blank"
  rel="noopener"><code>VkPhysicalDeviceAccelerationStructureFeaturesKHR</code>::<code>accelerationStructure</code></a>
  feature **must** be enabled

- <a
  href="#VUID-vkDestroyAccelerationStructureKHR-accelerationStructure-02442"
  id="VUID-vkDestroyAccelerationStructureKHR-accelerationStructure-02442"></a>
  VUID-vkDestroyAccelerationStructureKHR-accelerationStructure-02442  
  All submitted commands that refer to `accelerationStructure` **must**
  have completed execution

- <a
  href="#VUID-vkDestroyAccelerationStructureKHR-accelerationStructure-02443"
  id="VUID-vkDestroyAccelerationStructureKHR-accelerationStructure-02443"></a>
  VUID-vkDestroyAccelerationStructureKHR-accelerationStructure-02443  
  If `VkAllocationCallbacks` were provided when `accelerationStructure`
  was created, a compatible set of callbacks **must** be provided here

- <a
  href="#VUID-vkDestroyAccelerationStructureKHR-accelerationStructure-02444"
  id="VUID-vkDestroyAccelerationStructureKHR-accelerationStructure-02444"></a>
  VUID-vkDestroyAccelerationStructureKHR-accelerationStructure-02444  
  If no `VkAllocationCallbacks` were provided when
  `accelerationStructure` was created, `pAllocator` **must** be `NULL`

Valid Usage (Implicit)

- <a href="#VUID-vkDestroyAccelerationStructureKHR-device-parameter"
  id="VUID-vkDestroyAccelerationStructureKHR-device-parameter"></a>
  VUID-vkDestroyAccelerationStructureKHR-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a
  href="#VUID-vkDestroyAccelerationStructureKHR-accelerationStructure-parameter"
  id="VUID-vkDestroyAccelerationStructureKHR-accelerationStructure-parameter"></a>
  VUID-vkDestroyAccelerationStructureKHR-accelerationStructure-parameter  
  If `accelerationStructure` is not
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), `accelerationStructure`
  **must** be a valid
  [VkAccelerationStructureKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureKHR.html) handle

- <a href="#VUID-vkDestroyAccelerationStructureKHR-pAllocator-parameter"
  id="VUID-vkDestroyAccelerationStructureKHR-pAllocator-parameter"></a>
  VUID-vkDestroyAccelerationStructureKHR-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid
  pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html)
  structure

- <a
  href="#VUID-vkDestroyAccelerationStructureKHR-accelerationStructure-parent"
  id="VUID-vkDestroyAccelerationStructureKHR-accelerationStructure-parent"></a>
  VUID-vkDestroyAccelerationStructureKHR-accelerationStructure-parent  
  If `accelerationStructure` is a valid handle, it **must** have been
  created, allocated, or retrieved from `device`

Host Synchronization

- Host access to `accelerationStructure` **must** be externally
  synchronized

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_acceleration_structure](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_acceleration_structure.html),
[VkAccelerationStructureKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureKHR.html),
[VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkDestroyAccelerationStructureKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
