# LaunchIdKHR(3) Manual Page

## Name

LaunchIdKHR - Launch Id for ray shaders



## <a href="#_description" class="anchor"></a>Description

`LaunchIdKHR`  
A variable decorated with the `LaunchIdKHR` decoration will specify the
index of the work item being processed. One work item is generated for
each of the `width` × `height` × `depth` items dispatched by a
[vkCmdTraceRaysKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdTraceRaysKHR.html) command. All shader
invocations inherit the same value for variables decorated with
`LaunchIdKHR`.

Valid Usage

- <a href="#VUID-LaunchIdKHR-LaunchIdKHR-04266"
  id="VUID-LaunchIdKHR-LaunchIdKHR-04266"></a>
  VUID-LaunchIdKHR-LaunchIdKHR-04266  
  The `LaunchIdKHR` decoration **must** be used only within the
  `RayGenerationKHR`, `IntersectionKHR`, `AnyHitKHR`, `ClosestHitKHR`,
  `MissKHR`, or `CallableKHR` `Execution` `Model`

- <a href="#VUID-LaunchIdKHR-LaunchIdKHR-04267"
  id="VUID-LaunchIdKHR-LaunchIdKHR-04267"></a>
  VUID-LaunchIdKHR-LaunchIdKHR-04267  
  The variable decorated with `LaunchIdKHR` **must** be declared using
  the `Input` `Storage` `Class`

- <a href="#VUID-LaunchIdKHR-LaunchIdKHR-04268"
  id="VUID-LaunchIdKHR-LaunchIdKHR-04268"></a>
  VUID-LaunchIdKHR-LaunchIdKHR-04268  
  The variable decorated with `LaunchIdKHR` **must** be declared as a
  three-component vector of 32-bit integer values

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#LaunchIdKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
