# LaunchIdKHR(3) Manual Page

## Name

LaunchIdKHR - Launch Id for ray shaders



## [](#_description)Description

`LaunchIdKHR`

A variable decorated with the `LaunchIdKHR` decoration will specify the index of the work item being processed. One work item is generated for each of the `width` × `height` × `depth` items dispatched by a [vkCmdTraceRaysKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdTraceRaysKHR.html) or [vkCmdTraceRaysNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdTraceRaysNV.html) command. All shader invocations inherit the same value for variables decorated with `LaunchIdKHR`.

Valid Usage

- [](#VUID-LaunchIdKHR-LaunchIdKHR-04266)VUID-LaunchIdKHR-LaunchIdKHR-04266  
  The `LaunchIdKHR` decoration **must** be used only within the `RayGenerationKHR`, `IntersectionKHR`, `AnyHitKHR`, `ClosestHitKHR`, `MissKHR`, or `CallableKHR` `Execution` `Model`
- [](#VUID-LaunchIdKHR-LaunchIdKHR-04267)VUID-LaunchIdKHR-LaunchIdKHR-04267  
  The variable decorated with `LaunchIdKHR` **must** be declared using the `Input` `Storage` `Class`
- [](#VUID-LaunchIdKHR-LaunchIdKHR-04268)VUID-LaunchIdKHR-LaunchIdKHR-04268  
  The variable decorated with `LaunchIdKHR` **must** be declared as a three-component vector of 32-bit integer values

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#LaunchIdKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0