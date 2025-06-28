# SubgroupLocalInvocationId(3) Manual Page

## Name

SubgroupLocalInvocationId - ID of the invocation within a subgroup



## [](#_description)Description

`SubgroupLocalInvocationId`

Decorating a variable with the `SubgroupLocalInvocationId` builtin decoration will make that variable contain the index of the invocation within the subgroup. This variable is in range \[0,`SubgroupSize`-1].

If `VK_PIPELINE_SHADER_STAGE_CREATE_REQUIRE_FULL_SUBGROUPS_BIT` is specified, or if `module` declares SPIR-V version 1.6 or higher, and the local workgroup size in the X dimension of the `stage` is a multiple of [`SubgroupSize`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#interfaces-builtin-variables-sgs), full subgroups are enabled for that pipeline stage. When full subgroups are enabled, subgroups **must** be launched with all invocations active, i.e., there is an active invocation with `SubgroupLocalInvocationId` for each value in range \[0,`SubgroupSize`-1].

Note

There is no direct relationship between `SubgroupLocalInvocationId` and `LocalInvocationId` or `LocalInvocationIndex`. If the pipeline or shader object was created with full subgroups applications can compute their own local invocation index to serve the same purpose:

index = `SubgroupLocalInvocationId` + `SubgroupId` × `SubgroupSize`

If full subgroups are not enabled, some subgroups may be dispatched with inactive invocations that do not correspond to a local workgroup invocation, making the value of index unreliable.

Note

`VK_PIPELINE_SHADER_STAGE_CREATE_REQUIRE_FULL_SUBGROUPS_BIT` and `VK_SHADER_CREATE_REQUIRE_FULL_SUBGROUPS_BIT_EXT` are effectively deprecated when compiling SPIR-V 1.6 shaders, as this behavior is the default for Vulkan with SPIR-V 1.6. This is more aligned with developer expectations, and avoids applications unexpectedly breaking in the future.

Valid Usage

- [](#VUID-SubgroupLocalInvocationId-SubgroupLocalInvocationId-04380)VUID-SubgroupLocalInvocationId-SubgroupLocalInvocationId-04380  
  The variable decorated with `SubgroupLocalInvocationId` **must** be declared using the `Input` `Storage` `Class`
- [](#VUID-SubgroupLocalInvocationId-SubgroupLocalInvocationId-04381)VUID-SubgroupLocalInvocationId-SubgroupLocalInvocationId-04381  
  The variable decorated with `SubgroupLocalInvocationId` **must** be declared as a scalar 32-bit integer value

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#SubgroupLocalInvocationId)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0