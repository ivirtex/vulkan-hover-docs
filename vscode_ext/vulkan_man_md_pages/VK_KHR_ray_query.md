# VK_KHR_ray_query(3) Manual Page

## Name

VK_KHR_ray_query - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

349

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_KHR_spirv_1_4](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_spirv_1_4.html)  
and  
[VK_KHR_acceleration_structure](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_acceleration_structure.html)  

## <a href="#_spir_v_dependencies" class="anchor"></a>SPIR-V Dependencies

- [SPV_KHR_ray_query](https://htmlpreview.github.io/?https://github.com/KhronosGroup/SPIRV-Registry/blob/main/extensions/KHR/SPV_KHR_ray_query.html)

## <a href="#_contact" class="anchor"></a>Contact

- Daniel Koch <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_ray_query%5D%20@dgkoch%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_ray_query%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>dgkoch</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2020-11-12

**Interactions and External Dependencies**  
- This extension provides API support for
  [`GLSL_EXT_ray_query`](https://github.com/KhronosGroup/GLSL/blob/main/extensions/ext/GLSL_EXT_ray_query.txt)

**Contributors**  
- Matthäus Chajdas, AMD

- Greg Grebe, AMD

- Nicolai Hähnle, AMD

- Tobias Hector, AMD

- Dave Oldcorn, AMD

- Skyler Saleh, AMD

- Mathieu Robart, Arm

- Marius Bjorge, Arm

- Tom Olson, Arm

- Sebastian Tafuri, EA

- Henrik Rydgard, Embark

- Juan Cañada, Epic Games

- Patrick Kelly, Epic Games

- Yuriy O’Donnell, Epic Games

- Michael Doggett, Facebook/Oculus

- Andrew Garrard, Imagination

- Don Scorgie, Imagination

- Dae Kim, Imagination

- Joshua Barczak, Intel

- Slawek Grajewski, Intel

- Jeff Bolz, NVIDIA

- Pascal Gautron, NVIDIA

- Daniel Koch, NVIDIA

- Christoph Kubisch, NVIDIA

- Ashwin Lele, NVIDIA

- Robert Stepinski, NVIDIA

- Martin Stich, NVIDIA

- Nuno Subtil, NVIDIA

- Eric Werness, NVIDIA

- Jon Leech, Khronos

- Jeroen van Schijndel, OTOY

- Juul Joosten, OTOY

- Alex Bourd, Qualcomm

- Roman Larionov, Qualcomm

- David McAllister, Qualcomm

- Spencer Fricke, Samsung

- Lewis Gordon, Samsung

- Ralph Potter, Samsung

- Jasper Bekkers, Traverse Research

- Jesse Barker, Unity

- Baldur Karlsson, Valve

## <a href="#_description" class="anchor"></a>Description

Rasterization has been the dominant method to produce interactive
graphics, but increasing performance of graphics hardware has made ray
tracing a viable option for interactive rendering. Being able to
integrate ray tracing with traditional rasterization makes it easier for
applications to incrementally add ray traced effects to existing
applications or to do hybrid approaches with rasterization for primary
visibility and ray tracing for secondary queries.

Ray queries are available to all shader types, including graphics,
compute and ray tracing pipelines. Ray queries are not able to launch
additional shaders, instead returning traversal results to the calling
shader.

This extension adds support for the following SPIR-V extension in
Vulkan:

- `SPV_KHR_ray_query`

## <a href="#_new_structures" class="anchor"></a>New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html),
  [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html):

  - [VkPhysicalDeviceRayQueryFeaturesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceRayQueryFeaturesKHR.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_KHR_RAY_QUERY_EXTENSION_NAME`

- `VK_KHR_RAY_QUERY_SPEC_VERSION`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_RAY_QUERY_FEATURES_KHR`

## <a href="#_new_spir_v_capabilities" class="anchor"></a>New SPIR-V Capabilities

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#spirvenv-capabilities-table-RayQueryKHR"
  target="_blank" rel="noopener"><code>RayQueryKHR</code></a>

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#spirvenv-capabilities-table-RayTraversalPrimitiveCullingKHR"
  target="_blank"
  rel="noopener"><code>RayTraversalPrimitiveCullingKHR</code></a>

## <a href="#_sample_code" class="anchor"></a>Sample Code

Example of ray query in a GLSL shader, illustrating how to use ray
queries to determine whether a given position (at ray origin) is in
shadow or not, by tracing a ray towards the light, and checking for any
intersections with geometry occluding the light.

``` c
rayQueryEXT rq;

rayQueryInitializeEXT(rq, accStruct, gl_RayFlagsTerminateOnFirstHitEXT, cullMask, origin, tMin, direction, tMax);

// Traverse the acceleration structure and store information about the first intersection (if any)
rayQueryProceedEXT(rq);

if (rayQueryGetIntersectionTypeEXT(rq, true) == gl_RayQueryCommittedIntersectionNoneEXT) {
    // Not in shadow
}
```

## <a href="#_issues" class="anchor"></a>Issues

\(1\) What are the changes between the public provisional
(VK_KHR_ray_tracing v8) release and the final
(VK_KHR_acceleration_structure v11 / VK_KHR_ray_query v1) release?

- refactor VK_KHR_ray_tracing into 3 extensions, enabling implementation
  flexibility and decoupling ray query support from ray pipelines:

  - [`VK_KHR_acceleration_structure`](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_acceleration_structure.html)
    (for acceleration structure operations)

  - [`VK_KHR_ray_tracing_pipeline`](VK_KHR_ray_tracing_pipeline.html)
    (for ray tracing pipeline and shader stages)

  - [`VK_KHR_ray_query`](VK_KHR_ray_query.html) (for ray queries in
    existing shader stages)

- Update SPIRV capabilities to use `RayQueryKHR`

- extension is no longer provisional

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2020-11-12 (Mathieu Robart, Daniel Koch, Andrew Garrard)

  - Decomposition of the specification, from VK_KHR_ray_tracing to
    VK_KHR_ray_query (#1918,!3912)

  - update to use `RayQueryKHR` SPIR-V capability

  - add numerical limits for ray parameters (#2235,!3960)

  - relax formula for ray intersection candidate determination
    (#2322,!4080)

  - restrict traces to TLAS (#2239,!4141)

  - require `HitT` to be in ray interval for
    `OpRayQueryGenerateIntersectionKHR` (#2359,!4146)

  - add ray query shader stages for AS read bit (#2407,!4203)

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_KHR_ray_query"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
