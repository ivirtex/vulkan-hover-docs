# VkStridedDeviceAddressRegionKHR(3) Manual Page

## Name

VkStridedDeviceAddressRegionKHR - Structure specifying a region of
device addresses with a stride



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkStridedDeviceAddressRegionKHR` structure is defined as:

``` c
// Provided by VK_KHR_ray_tracing_pipeline
typedef struct VkStridedDeviceAddressRegionKHR {
    VkDeviceAddress    deviceAddress;
    VkDeviceSize       stride;
    VkDeviceSize       size;
} VkStridedDeviceAddressRegionKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `deviceAddress` is the device address (as returned by the
  [vkGetBufferDeviceAddress](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetBufferDeviceAddress.html) command) at
  which the region starts, or zero if the region is unused.

- `stride` is the byte stride between consecutive elements.

- `size` is the size in bytes of the region starting at `deviceAddress`.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-VkStridedDeviceAddressRegionKHR-size-04631"
  id="VUID-VkStridedDeviceAddressRegionKHR-size-04631"></a>
  VUID-VkStridedDeviceAddressRegionKHR-size-04631  
  If `size` is not zero, all addresses between `deviceAddress` and
  `deviceAddress` + `size` - 1 **must** be in the buffer device address
  range of the same buffer

- <a href="#VUID-VkStridedDeviceAddressRegionKHR-size-04632"
  id="VUID-VkStridedDeviceAddressRegionKHR-size-04632"></a>
  VUID-VkStridedDeviceAddressRegionKHR-size-04632  
  If `size` is not zero, `stride` **must** be less than or equal to the
  size of the buffer from which `deviceAddress` was queried

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_ray_tracing_pipeline](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_ray_tracing_pipeline.html),
[VkDeviceAddress](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceAddress.html),
[VkDeviceSize](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceSize.html),
[vkCmdTraceRaysIndirectKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdTraceRaysIndirectKHR.html),
[vkCmdTraceRaysKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdTraceRaysKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkStridedDeviceAddressRegionKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
