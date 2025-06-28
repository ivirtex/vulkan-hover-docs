# VK\_KHR\_ray\_query(3) Manual Page

## Name

VK\_KHR\_ray\_query - device extension



## [](#_registered_extension_number)Registered Extension Number

349

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

     [VK\_KHR\_spirv\_1\_4](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_spirv_1_4.html)  
     or  
     [Vulkan Version 1.2](#versions-1.2)  
and  
[VK\_KHR\_acceleration\_structure](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_acceleration_structure.html)

## [](#_spir_v_dependencies)SPIR-V Dependencies

- [SPV\_KHR\_ray\_query](https://github.khronos.org/SPIRV-Registry/extensions/KHR/SPV_KHR_ray_query.html)

## [](#_contact)Contact

- Daniel Koch [\[GitHub\]dgkoch](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_ray_query%5D%20%40dgkoch%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_ray_query%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2020-11-12

**Interactions and External Dependencies**

- This extension provides API support for [`GLSL_EXT_ray_query`](https://github.com/KhronosGroup/GLSL/blob/main/extensions/ext/GLSL_EXT_ray_query.txt)

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

## [](#_description)Description

Rasterization has been the dominant method to produce interactive graphics, but increasing performance of graphics hardware has made ray tracing a viable option for interactive rendering. Being able to integrate ray tracing with traditional rasterization makes it easier for applications to incrementally add ray traced effects to existing applications or to do hybrid approaches with rasterization for primary visibility and ray tracing for secondary queries.

Ray queries are available to all shader types, including graphics, compute and ray tracing pipelines. Ray queries are not able to launch additional shaders, instead returning traversal results to the calling shader.

This extension adds support for the following SPIR-V extension in Vulkan:

- `SPV_KHR_ray_query`

## [](#_new_structures)New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceRayQueryFeaturesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceRayQueryFeaturesKHR.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_KHR_RAY_QUERY_EXTENSION_NAME`
- `VK_KHR_RAY_QUERY_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_RAY_QUERY_FEATURES_KHR`

## [](#_new_spir_v_capabilities)New SPIR-V Capabilities

- [`RayQueryKHR`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#spirvenv-capabilities-table-RayQueryKHR)
- [`RayTraversalPrimitiveCullingKHR`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#spirvenv-capabilities-table-RayTraversalPrimitiveCullingKHR)

## [](#_sample_code)Sample Code

Example of ray query in a GLSL shader, illustrating how to use ray queries to determine whether a given position (at ray origin) is in shadow or not, by tracing a ray towards the light, and checking for any intersections with geometry occluding the light.

```c
rayQueryEXT rq;

rayQueryInitializeEXT(rq, accStruct, gl_RayFlagsTerminateOnFirstHitEXT, cullMask, origin, tMin, direction, tMax);

// Traverse the acceleration structure and store information about the first intersection (if any)
rayQueryProceedEXT(rq);

if (rayQueryGetIntersectionTypeEXT(rq, true) == gl_RayQueryCommittedIntersectionNoneEXT) {
    // Not in shadow
}
```

## [](#_issues)Issues

(1) What are the changes between the public provisional (VK\_KHR\_ray\_tracing v8) release and the final (VK\_KHR\_acceleration\_structure v11 / VK\_KHR\_ray\_query v1) release?

- refactor VK\_KHR\_ray\_tracing into 3 extensions, enabling implementation flexibility and decoupling ray query support from ray pipelines:
  
  - `VK_KHR_acceleration_structure` (for acceleration structure operations)
  - `VK_KHR_ray_tracing_pipeline` (for ray tracing pipeline and shader stages)
  - `VK_KHR_ray_query` (for ray queries in existing shader stages)
- Update SPIRV capabilities to use `RayQueryKHR`
- extension is no longer provisional

## [](#_version_history)Version History

- Revision 1, 2020-11-12 (Mathieu Robart, Daniel Koch, Andrew Garrard)
  
  - Decomposition of the specification, from VK\_KHR\_ray\_tracing to VK\_KHR\_ray\_query (#1918,!3912)
  - update to use `RayQueryKHR` SPIR-V capability
  - add numerical limits for ray parameters (#2235,!3960)
  - relax formula for ray intersection candidate determination (#2322,!4080)
  - restrict traces to TLAS (#2239,!4141)
  - require `HitT` to be in ray interval for `OpRayQueryGenerateIntersectionKHR` (#2359,!4146)
  - add ray query shader stages for AS read bit (#2407,!4203)

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_KHR_ray_query)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0