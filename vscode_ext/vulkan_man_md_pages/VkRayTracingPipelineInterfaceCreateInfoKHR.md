# VkRayTracingPipelineInterfaceCreateInfoKHR(3) Manual Page

## Name

VkRayTracingPipelineInterfaceCreateInfoKHR - Structure specifying
additional interface information when using libraries



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkRayTracingPipelineInterfaceCreateInfoKHR` structure is defined
as:

``` c
// Provided by VK_KHR_ray_tracing_pipeline
typedef struct VkRayTracingPipelineInterfaceCreateInfoKHR {
    VkStructureType    sType;
    const void*        pNext;
    uint32_t           maxPipelineRayPayloadSize;
    uint32_t           maxPipelineRayHitAttributeSize;
} VkRayTracingPipelineInterfaceCreateInfoKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `maxPipelineRayPayloadSize` is the maximum payload size in bytes used
  by any shader in the pipeline.

- `maxPipelineRayHitAttributeSize` is the maximum attribute structure
  size in bytes used by any shader in the pipeline.

## <a href="#_description" class="anchor"></a>Description

`maxPipelineRayPayloadSize` is calculated as the maximum number of bytes
used by any block declared in the `RayPayloadKHR` or
`IncomingRayPayloadKHR` storage classes.
`maxPipelineRayHitAttributeSize` is calculated as the maximum number of
bytes used by any block declared in the `HitAttributeKHR` storage class.
As variables in these storage classes do not have explicit offsets, the
size should be calculated as if each variable has a <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#interfaces-alignment-requirements"
target="_blank" rel="noopener">scalar alignment</a> equal to the largest
scalar alignment of any of the block’s members.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr class="odd">
<td class="icon"><em></em></td>
<td class="content">Note
<p>There is no explicit upper limit for
<code>maxPipelineRayPayloadSize</code>, but in practice it should be
kept as small as possible. Similar to invocation local memory, it must
be allocated for each shader invocation and for devices which support
many simultaneous invocations, this storage can rapidly be exhausted,
resulting in failure.</p></td>
</tr>
</tbody>
</table>

Valid Usage

- <a
  href="#VUID-VkRayTracingPipelineInterfaceCreateInfoKHR-maxPipelineRayHitAttributeSize-03605"
  id="VUID-VkRayTracingPipelineInterfaceCreateInfoKHR-maxPipelineRayHitAttributeSize-03605"></a>
  VUID-VkRayTracingPipelineInterfaceCreateInfoKHR-maxPipelineRayHitAttributeSize-03605  
  `maxPipelineRayHitAttributeSize` **must** be less than or equal to
  [VkPhysicalDeviceRayTracingPipelinePropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceRayTracingPipelinePropertiesKHR.html)::`maxRayHitAttributeSize`

Valid Usage (Implicit)

- <a href="#VUID-VkRayTracingPipelineInterfaceCreateInfoKHR-sType-sType"
  id="VUID-VkRayTracingPipelineInterfaceCreateInfoKHR-sType-sType"></a>
  VUID-VkRayTracingPipelineInterfaceCreateInfoKHR-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_RAY_TRACING_PIPELINE_INTERFACE_CREATE_INFO_KHR`

- <a href="#VUID-VkRayTracingPipelineInterfaceCreateInfoKHR-pNext-pNext"
  id="VUID-VkRayTracingPipelineInterfaceCreateInfoKHR-pNext-pNext"></a>
  VUID-VkRayTracingPipelineInterfaceCreateInfoKHR-pNext-pNext  
  `pNext` **must** be `NULL`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_ray_tracing_pipeline](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_ray_tracing_pipeline.html),
[VkRayTracingPipelineCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRayTracingPipelineCreateInfoKHR.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkRayTracingPipelineInterfaceCreateInfoKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
