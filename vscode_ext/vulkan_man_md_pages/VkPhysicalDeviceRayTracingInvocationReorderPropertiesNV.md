# VkPhysicalDeviceRayTracingInvocationReorderPropertiesNV(3) Manual Page

## Name

VkPhysicalDeviceRayTracingInvocationReorderPropertiesNV - Structure
describing shader module identifier properties of an implementation



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceRayTracingInvocationReorderPropertiesNV` structure
is defined as:

``` c
// Provided by VK_NV_ray_tracing_invocation_reorder
typedef struct VkPhysicalDeviceRayTracingInvocationReorderPropertiesNV {
    VkStructureType                        sType;
    void*                                  pNext;
    VkRayTracingInvocationReorderModeNV    rayTracingInvocationReorderReorderingHint;
} VkPhysicalDeviceRayTracingInvocationReorderPropertiesNV;
```

## <a href="#_members" class="anchor"></a>Members

- `rayTracingInvocationReorderReorderingHint` is a hint indicating if
  the implementation will actually reorder at the reorder calls.

## <a href="#_description" class="anchor"></a>Description

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr>
<td class="icon"><em></em></td>
<td class="content">Note
<p>Because the extension changes how hits are managed there is a
compatibility reason to expose the extension even when an implementation
does not have sorting active.</p></td>
</tr>
</tbody>
</table>

If the `VkPhysicalDeviceRayTracingInvocationReorderPropertiesNV`
structure is included in the `pNext` chain of the
[VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties2.html)
structure passed to
[vkGetPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceProperties2.html),
it is filled in with each corresponding implementation-dependent
property.

Valid Usage (Implicit)

- <a
  href="#VUID-VkPhysicalDeviceRayTracingInvocationReorderPropertiesNV-sType-sType"
  id="VUID-VkPhysicalDeviceRayTracingInvocationReorderPropertiesNV-sType-sType"></a>
  VUID-VkPhysicalDeviceRayTracingInvocationReorderPropertiesNV-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_RAY_TRACING_INVOCATION_REORDER_PROPERTIES_NV`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_ray_tracing_invocation_reorder](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_ray_tracing_invocation_reorder.html),
[VkRayTracingInvocationReorderModeNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRayTracingInvocationReorderModeNV.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceRayTracingInvocationReorderPropertiesNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
