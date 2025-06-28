# StandaloneSpirv(3) Manual Page

## Name

StandaloneSpirv - Standalone SPIR-V Validation



## [](#_description)Description

The following rules **can** be validated with only the SPIR-V module itself. They do not depend on knowledge of the implementation and its capabilities or knowledge of runtime information, such as enabled features.

Valid Usage

- [](#VUID-StandaloneSpirv-None-04633)VUID-StandaloneSpirv-None-04633  
  Every entry point **must** have no return value and accept no arguments
- [](#VUID-StandaloneSpirv-None-04634)VUID-StandaloneSpirv-None-04634  
  The static function-call graph for an entry point **must** not contain cycles; that is, static recursion is not allowed
- [](#VUID-StandaloneSpirv-None-04635)VUID-StandaloneSpirv-None-04635  
  The `Logical` or `PhysicalStorageBuffer64` addressing model **must** be selected
- [](#VUID-StandaloneSpirv-None-04636)VUID-StandaloneSpirv-None-04636  
  `Scope` for execution **must** be limited to `Workgroup` or `Subgroup`
- [](#VUID-StandaloneSpirv-None-04637)VUID-StandaloneSpirv-None-04637  
  If the `Scope` for execution is `Workgroup`, then it **must** only be used in the task, mesh, tessellation control, or compute `Execution` `Model`
- [](#VUID-StandaloneSpirv-None-04638)VUID-StandaloneSpirv-None-04638  
  `Scope` for memory **must** be limited to `Device`, `QueueFamily`, `Workgroup`, `ShaderCallKHR`, `Subgroup`, or `Invocation`
- [](#VUID-StandaloneSpirv-ExecutionModel-07320)VUID-StandaloneSpirv-ExecutionModel-07320  
  If the `Execution` `Model` is `TessellationControl`, and the `MemoryModel` is `GLSL450`, the `Scope` for memory **must** not be `Workgroup`
- [](#VUID-StandaloneSpirv-None-07321)VUID-StandaloneSpirv-None-07321  
  If the `Scope` for memory is `Workgroup`, then it **must** only be used in the task, mesh, tessellation control, or compute `Execution` `Model`
- [](#VUID-StandaloneSpirv-None-04640)VUID-StandaloneSpirv-None-04640  
  If the `Scope` for memory is `ShaderCallKHR`, then it **must** only be used in ray generation, intersection, closest hit, any-hit, miss, and callable `Execution` `Model`
- [](#VUID-StandaloneSpirv-None-04641)VUID-StandaloneSpirv-None-04641  
  If the `Scope` for memory is `Invocation`, then memory semantics **must** be `None`
- [](#VUID-StandaloneSpirv-None-04642)VUID-StandaloneSpirv-None-04642  
  `Scope` for [group operations](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#shaders-group-operations) **must** be limited to `Subgroup`
- [](#VUID-StandaloneSpirv-SubgroupVoteKHR-07951)VUID-StandaloneSpirv-SubgroupVoteKHR-07951  
  If none of the `SubgroupVoteKHR`, `GroupNonUniform`, or `SubgroupBallotKHR` capabilities are declared, `Scope` for memory **must** not be `Subgroup`
- [](#VUID-StandaloneSpirv-None-04643)VUID-StandaloneSpirv-None-04643  
  `Storage` `Class` **must** be limited to `UniformConstant`, `Input`, `Uniform`, `Output`, `Workgroup`, `Private`, `Function`, `PushConstant`, `Image`, `StorageBuffer`, `RayPayloadKHR`, `IncomingRayPayloadKHR`, `HitAttributeKHR`, `CallableDataKHR`, `IncomingCallableDataKHR`, `ShaderRecordBufferKHR`, `PhysicalStorageBuffer`, or `TileImageEXT`
- [](#VUID-StandaloneSpirv-None-04644)VUID-StandaloneSpirv-None-04644  
  If the `Storage` `Class` is `Output`, then it **must** not be used in the `GlCompute`, `RayGenerationKHR`, `IntersectionKHR`, `AnyHitKHR`, `ClosestHitKHR`, `MissKHR`, or `CallableKHR` `Execution` `Model`
- [](#VUID-StandaloneSpirv-None-04645)VUID-StandaloneSpirv-None-04645  
  If the `Storage` `Class` is `Workgroup`, then it **must** only be used in the task, mesh, or compute `Execution` `Model`
- [](#VUID-StandaloneSpirv-None-08720)VUID-StandaloneSpirv-None-08720  
  If the `Storage` `Class` is `TileImageEXT`, then it **must** only be used in the fragment execution model
- [](#VUID-StandaloneSpirv-OpAtomicStore-04730)VUID-StandaloneSpirv-OpAtomicStore-04730  
  `OpAtomicStore` **must** not use `Acquire`, `AcquireRelease`, or `SequentiallyConsistent` memory semantics
- [](#VUID-StandaloneSpirv-OpAtomicLoad-04731)VUID-StandaloneSpirv-OpAtomicLoad-04731  
  `OpAtomicLoad` **must** not use `Release`, `AcquireRelease`, or `SequentiallyConsistent` memory semantics
- [](#VUID-StandaloneSpirv-OpMemoryBarrier-04732)VUID-StandaloneSpirv-OpMemoryBarrier-04732  
  `OpMemoryBarrier` **must** use one of `Acquire`, `Release`, `AcquireRelease`, or `SequentiallyConsistent` memory semantics
- [](#VUID-StandaloneSpirv-OpMemoryBarrier-04733)VUID-StandaloneSpirv-OpMemoryBarrier-04733  
  `OpMemoryBarrier` **must** include at least one `Storage` `Class`
- [](#VUID-StandaloneSpirv-OpControlBarrier-10609)VUID-StandaloneSpirv-OpControlBarrier-10609  
  If the semantics for `OpControlBarrier` includes at least one `Storage` `Class`, then it **must** include one of `Acquire`, `Release`, `AcquireRelease`, or `SequentiallyConsistent` memory semantics
- [](#VUID-StandaloneSpirv-OpControlBarrier-04650)VUID-StandaloneSpirv-OpControlBarrier-04650  
  If the semantics for `OpControlBarrier` includes one of `Acquire`, `Release`, `AcquireRelease`, or `SequentiallyConsistent` memory semantics, then it **must** include at least one `Storage` `Class`
- [](#VUID-StandaloneSpirv-OpVariable-04651)VUID-StandaloneSpirv-OpVariable-04651  
  Any `OpVariable` with an `Initializer` operand **must** have `Output`, `Private`, `Function`, or `Workgroup` as its `Storage` `Class` operand
- [](#VUID-StandaloneSpirv-OpVariable-04734)VUID-StandaloneSpirv-OpVariable-04734  
  Any `OpVariable` with an `Initializer` operand and `Workgroup` as its `Storage` `Class` operand **must** use `OpConstantNull` as the initializer
- [](#VUID-StandaloneSpirv-OpReadClockKHR-04652)VUID-StandaloneSpirv-OpReadClockKHR-04652  
  `Scope` for `OpReadClockKHR` **must** be limited to `Subgroup` or `Device`
- [](#VUID-StandaloneSpirv-OriginLowerLeft-04653)VUID-StandaloneSpirv-OriginLowerLeft-04653  
  The `OriginLowerLeft` `Execution` `Mode` **must** not be used; fragment entry points **must** declare `OriginUpperLeft`
- [](#VUID-StandaloneSpirv-PixelCenterInteger-04654)VUID-StandaloneSpirv-PixelCenterInteger-04654  
  The `PixelCenterInteger` `Execution` `Mode` **must** not be used (pixels are always centered at half-integer coordinates)
- [](#VUID-StandaloneSpirv-UniformConstant-04655)VUID-StandaloneSpirv-UniformConstant-04655  
  Any variable in the `UniformConstant` `Storage` `Class` **must** be typed as either `OpTypeImage`, `OpTypeSampler`, `OpTypeSampledImage`, `OpTypeAccelerationStructureKHR`, `OpTypeTensorARM` or an array of one of these types
- [](#VUID-StandaloneSpirv-Uniform-06807)VUID-StandaloneSpirv-Uniform-06807  
  Any variable in the `Uniform` or `StorageBuffer` `Storage` `Class` **must** be typed as `OpTypeStruct` or an array of this type
- [](#VUID-StandaloneSpirv-PushConstant-06808)VUID-StandaloneSpirv-PushConstant-06808  
  Any variable in the `PushConstant` `Storage` `Class` **must** be typed as `OpTypeStruct`
- [](#VUID-StandaloneSpirv-OpTypeImage-04656)VUID-StandaloneSpirv-OpTypeImage-04656  
  `OpTypeImage` **must** declare a scalar 32-bit float, 64-bit integer, or 32-bit integer type for the “Sampled Type” (`RelaxedPrecision` **can** be applied to a sampling instruction and to the variable holding the result of a sampling instruction)
- [](#VUID-StandaloneSpirv-OpTypeImage-04657)VUID-StandaloneSpirv-OpTypeImage-04657  
  `OpTypeImage` **must** have a “Sampled” operand of 1 (sampled image) or 2 (storage image)
- [](#VUID-StandaloneSpirv-OpTypeSampledImage-06671)VUID-StandaloneSpirv-OpTypeSampledImage-06671  
  `OpTypeSampledImage` **must** have a `OpTypeImage` with a “Sampled” operand of 1 (sampled image)
- [](#VUID-StandaloneSpirv-Image-04965)VUID-StandaloneSpirv-Image-04965  
  The [SPIR-V Type](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#spirv-type) of the `Image` `Format` operand of an `OpTypeImage` **must** match the `Sampled` `Type`, as defined in [https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#spirvenv-format-type-matching](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#spirvenv-format-type-matching)
- [](#VUID-StandaloneSpirv-OpImageTexelPointer-04658)VUID-StandaloneSpirv-OpImageTexelPointer-04658  
  If an `OpImageTexelPointer` is used in an atomic operation, the image type of the `image` parameter to `OpImageTexelPointer` **must** have an image format of `R64i`, `R64ui`, `R32f`, `R32i`, or `R32ui`
- [](#VUID-StandaloneSpirv-OpImageQuerySizeLod-04659)VUID-StandaloneSpirv-OpImageQuerySizeLod-04659  
  `OpImageQuerySizeLod`, `OpImageQueryLod`, and `OpImageQueryLevels` **must** only consume an “Image” operand whose type has its “Sampled” operand set to 1
- [](#VUID-StandaloneSpirv-OpTypeImage-09638)VUID-StandaloneSpirv-OpTypeImage-09638  
  An `OpTypeImage` **must** not have a “Dim” operand of `Rect`
- [](#VUID-StandaloneSpirv-OpTypeImage-06214)VUID-StandaloneSpirv-OpTypeImage-06214  
  An `OpTypeImage` with a “Dim” operand of `SubpassData` **must** have an “Arrayed” operand of 0 (non-arrayed) and a “Sampled” operand of 2 (storage image)
- [](#VUID-StandaloneSpirv-SubpassData-04660)VUID-StandaloneSpirv-SubpassData-04660  
  The (u,v) coordinates used for a `SubpassData` **must** be the &lt;id&gt; of a constant vector (0,0)
- [](#VUID-StandaloneSpirv-OpTypeImage-06924)VUID-StandaloneSpirv-OpTypeImage-06924  
  Objects of types `OpTypeImage`, `OpTypeSampler`, `OpTypeSampledImage`, `OpTypeAccelerationStructureKHR`, `OpTypeTensorARM`, and arrays of these types **must** not be stored to or modified
- [](#VUID-StandaloneSpirv-Uniform-06925)VUID-StandaloneSpirv-Uniform-06925  
  Any variable in the `Uniform` `Storage` `Class` decorated as `Block` **must** not be stored to or modified
- [](#VUID-StandaloneSpirv-Offset-04865)VUID-StandaloneSpirv-Offset-04865  
  Any image instruction which uses an `Offset`, `ConstOffset`, or `ConstOffsets` image operand, **must** only consume a “Sampled Image” operand whose type has its “Sampled” operand set to 1
- [](#VUID-StandaloneSpirv-OpImageGather-04664)VUID-StandaloneSpirv-OpImageGather-04664  
  The “Component” operand of `OpImageGather`, and `OpImageSparseGather` **must** be the &lt;id&gt; of a constant instruction
- [](#VUID-StandaloneSpirv-OpImage-04777)VUID-StandaloneSpirv-OpImage-04777  
  `OpImage*Dref*` instructions **must** not consume an image whose `Dim` is 3D
- [](#VUID-StandaloneSpirv-None-04667)VUID-StandaloneSpirv-None-04667  
  Structure types **must** not contain opaque types
- [](#VUID-StandaloneSpirv-BuiltIn-04668)VUID-StandaloneSpirv-BuiltIn-04668  
  Any `BuiltIn` decoration not listed in [https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#interfaces-builtin-variables](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#interfaces-builtin-variables) **must** not be used
- [](#VUID-StandaloneSpirv-OpEntryPoint-09658)VUID-StandaloneSpirv-OpEntryPoint-09658  
  For a given `OpEntryPoint`, any `BuiltIn` decoration **must** not be used more than once by the `Input` interface
- [](#VUID-StandaloneSpirv-OpEntryPoint-09659)VUID-StandaloneSpirv-OpEntryPoint-09659  
  For a given `OpEntryPoint`, any `BuiltIn` decoration **must** not be used more than once by the `Output` interface
- [](#VUID-StandaloneSpirv-Location-06672)VUID-StandaloneSpirv-Location-06672  
  The `Location` or `Component` decorations **must** only be used with the `Input`, `Output`, `RayPayloadKHR`, `IncomingRayPayloadKHR`, `HitAttributeKHR`, `HitObjectAttributeNV`, `CallableDataKHR`, `IncomingCallableDataKHR`, or `ShaderRecordBufferKHR` storage classes
- [](#VUID-StandaloneSpirv-Location-04915)VUID-StandaloneSpirv-Location-04915  
  The `Location` or `Component` decorations **must** not be used with `BuiltIn`
- [](#VUID-StandaloneSpirv-Location-04916)VUID-StandaloneSpirv-Location-04916  
  The `Location` decorations **must** be used on [user-defined variables](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#interfaces-iointerfaces-user)
- [](#VUID-StandaloneSpirv-Location-04917)VUID-StandaloneSpirv-Location-04917  
  If a [user-defined variable](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#interfaces-iointerfaces-user) is not a pointer to a `Block` decorated `OpTypeStruct`, then the `OpVariable` **must** have a `Location` decoration
- [](#VUID-StandaloneSpirv-Location-04918)VUID-StandaloneSpirv-Location-04918  
  If a [user-defined variable](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#interfaces-iointerfaces-user) has a `Location` decoration, and the variable is a pointer to a `OpTypeStruct`, then the members of that structure **must** not have `Location` decorations
- [](#VUID-StandaloneSpirv-Location-04919)VUID-StandaloneSpirv-Location-04919  
  If a [user-defined variable](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#interfaces-iointerfaces-user) does not have a `Location` decoration, and the variable is a pointer to a `Block` decorated `OpTypeStruct`, then each member of the struct **must** have a `Location` decoration
- [](#VUID-StandaloneSpirv-Component-04920)VUID-StandaloneSpirv-Component-04920  
  The `Component` decoration value **must** not be greater than 3
- [](#VUID-StandaloneSpirv-Component-04921)VUID-StandaloneSpirv-Component-04921  
  If the `Component` decoration is used on an `OpVariable` that has a `OpTypeVector` type with a `Component` `Type` with a `Width` that is less than or equal to 32, the sum of its `Component` `Count` and the `Component` decoration value **must** be less than or equal to 4
- [](#VUID-StandaloneSpirv-Component-04922)VUID-StandaloneSpirv-Component-04922  
  If the `Component` decoration is used on an `OpVariable` that has a `OpTypeVector` type with a `Component` `Type` with a `Width` that is equal to 64, the sum of two times its `Component` `Count` and the `Component` decoration value **must** be less than or equal to 4
- [](#VUID-StandaloneSpirv-Component-04923)VUID-StandaloneSpirv-Component-04923  
  The `Component` decorations value **must** not be 1 or 3 for scalar or two-component 64-bit data types
- [](#VUID-StandaloneSpirv-Component-07703)VUID-StandaloneSpirv-Component-07703  
  The `Component` decorations **must** not be used for a 64-bit vector type with more than two components
- [](#VUID-StandaloneSpirv-Component-10583)VUID-StandaloneSpirv-Component-10583  
  The `Component` decorations **must** not be used with any type that is not a scalar, vector, array of scalars or vectors, or an array of arrays of scalars or vectors
- [](#VUID-StandaloneSpirv-Execution-10584)VUID-StandaloneSpirv-Execution-10584  
  If the `Execution` `Model` of an entry point is not `TessellationControl` or `Geometry`, its interface definition **must** not include any variables in the `Input` storage class decorated with `Component` that are arrays of arrays of scalars or vectors
- [](#VUID-StandaloneSpirv-Execution-10585)VUID-StandaloneSpirv-Execution-10585  
  If the `Execution` `Model` of an entry point is not `MeshNV`, its interface definition **must** not include any variables in the `Output` storage class decorated with `Component` that are arrays of arrays of scalars or vectors
- [](#VUID-StandaloneSpirv-Output-10586)VUID-StandaloneSpirv-Output-10586  
  Variables in the `Output` storage class in the `Vertex`, `TessellationControl`, `TessellationEvaluation`, or `Geometry` execution model **must** not have overlapping `Component` and `Location` decorations as defined by [Location and Component Assignment](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#interfaces-iointerfaces-locations)
- [](#VUID-StandaloneSpirv-Output-10587)VUID-StandaloneSpirv-Output-10587  
  Variables in the `Output` storage class in the `Fragment` execution model **must** not have both identical `Index` decorations and overlapping `Component` and `Location` decorations as defined by [Location and Component Assignment](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#interfaces-iointerfaces-locations)
- [](#VUID-StandaloneSpirv-Input-10588)VUID-StandaloneSpirv-Input-10588  
  Variables in the `Input` storage class in the `TessellationControl`, `TessellationEvaluation`, `Geometry`, or `Fragment` execution model **must** not have overlapping `Component` and `Location` decorations as defined by [Location and Component Assignment](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#interfaces-iointerfaces-locations)
- [](#VUID-StandaloneSpirv-Input-10604)VUID-StandaloneSpirv-Input-10604  
  Any variable in the `Input` storage class in the `Fragment` execution model **must** not have an assigned `Location` shared with another variable with different decorations for `PerVertexKHR`, `Flat`, `NoPerspective`, `Sample`, or `Centroid`
- [](#VUID-StandaloneSpirv-Input-09557)VUID-StandaloneSpirv-Input-09557  
  The pointers of any `Input` or `Output` [Interface user-defined variables](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#interfaces-iointerfaces-user) **must** not contain any `PhysicalStorageBuffer` `Storage` `Class` pointers
- [](#VUID-StandaloneSpirv-None-10684)VUID-StandaloneSpirv-None-10684  
  All variables **must** have valid explicit layout decorations [as described in Shader Interfaces](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#interfaces-explicit-layout-decorations)
- [](#VUID-StandaloneSpirv-GLSLShared-04669)VUID-StandaloneSpirv-GLSLShared-04669  
  The `GLSLShared` and `GLSLPacked` decorations **must** not be used
- [](#VUID-StandaloneSpirv-Flat-04670)VUID-StandaloneSpirv-Flat-04670  
  The `Flat`, `NoPerspective`, `Sample`, and `Centroid` decorations **must** only be used on variables with the `Output` or `Input` `Storage` `Class`
- [](#VUID-StandaloneSpirv-Flat-06201)VUID-StandaloneSpirv-Flat-06201  
  The `Flat`, `NoPerspective`, `Sample`, and `Centroid` decorations **must** not be used on variables with the `Output` storage class in a fragment shader
- [](#VUID-StandaloneSpirv-Flat-06202)VUID-StandaloneSpirv-Flat-06202  
  The `Flat`, `NoPerspective`, `Sample`, and `Centroid` decorations **must** not be used on variables with the `Input` storage class in a vertex shader
- [](#VUID-StandaloneSpirv-PerVertexKHR-06777)VUID-StandaloneSpirv-PerVertexKHR-06777  
  The `PerVertexKHR` decoration **must** only be used on variables with the `Input` `Storage` `Class` in a fragment shader
- [](#VUID-StandaloneSpirv-Flat-04744)VUID-StandaloneSpirv-Flat-04744  
  Any variable with integer or double-precision floating-point type and with `Input` `Storage` `Class` in a fragment shader, **must** be decorated `Flat`
- [](#VUID-StandaloneSpirv-ViewportRelativeNV-04672)VUID-StandaloneSpirv-ViewportRelativeNV-04672  
  The `ViewportRelativeNV` decoration **must** only be used on a variable decorated with `Layer` in the vertex, tessellation evaluation, or geometry shader stages
- [](#VUID-StandaloneSpirv-ViewportRelativeNV-04673)VUID-StandaloneSpirv-ViewportRelativeNV-04673  
  The `ViewportRelativeNV` decoration **must** not be used unless a variable decorated with one of `ViewportIndex` or `ViewportMaskNV` is also statically used by the same `OpEntryPoint`
- [](#VUID-StandaloneSpirv-ViewportMaskNV-04674)VUID-StandaloneSpirv-ViewportMaskNV-04674  
  The `ViewportMaskNV` and `ViewportIndex` decorations **must** not both be statically used by one or more `OpEntryPoint`’s that form the [pre-rasterization shader stages](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#pipelines-graphics-subsets-pre-rasterization) of a graphics pipeline
- [](#VUID-StandaloneSpirv-FPRoundingMode-04675)VUID-StandaloneSpirv-FPRoundingMode-04675  
  Rounding modes other than round-to-nearest-even and round-towards-zero **must** not be used for the `FPRoundingMode` decoration
- [](#VUID-StandaloneSpirv-Invariant-04677)VUID-StandaloneSpirv-Invariant-04677  
  Variables decorated with `Invariant` and variables with structure types that have any members decorated with `Invariant` **must** be in the `Output` or `Input` `Storage` `Class`, `Invariant` used on an `Input` `Storage` `Class` variable or structure member has no effect
- [](#VUID-StandaloneSpirv-VulkanMemoryModel-04678)VUID-StandaloneSpirv-VulkanMemoryModel-04678  
  []()If the `VulkanMemoryModel` capability is not declared, the `Volatile` decoration **must** be used on any variable declaration that includes one of the `SMIDNV`, `WarpIDNV`, `SubgroupSize`, `SubgroupLocalInvocationId`, `SubgroupEqMask`, `SubgroupGeMask`, `SubgroupGtMask`, `SubgroupLeMask`, or `SubgroupLtMask` `BuiltIn` decorations when used in the ray generation, closest hit, miss, intersection, or callable shaders, or with the `RayTmaxKHR` `Builtin` decoration when used in an intersection shader
- [](#VUID-StandaloneSpirv-VulkanMemoryModel-04679)VUID-StandaloneSpirv-VulkanMemoryModel-04679  
  If the `VulkanMemoryModel` capability is declared, the `OpLoad` instruction **must** use the `Volatile` memory semantics when it accesses into any variable that includes one of the `SMIDNV`, `WarpIDNV`, `SubgroupSize`, `SubgroupLocalInvocationId`, `SubgroupEqMask`, `SubgroupGeMask`, `SubgroupGtMask`, `SubgroupLeMask`, or `SubgroupLtMask` `BuiltIn` decorations when used in the ray generation, closest hit, miss, intersection, or callable shaders, or with the `RayTmaxKHR` `Builtin` decoration when used in an intersection shader
- [](#VUID-StandaloneSpirv-OpTypeRuntimeArray-04680)VUID-StandaloneSpirv-OpTypeRuntimeArray-04680  
  `OpTypeRuntimeArray` **must** only be used for:
  
  - the last member of a `Block`-decorated `OpTypeStruct` in `StorageBuffer` or `PhysicalStorageBuffer` storage `Storage` `Class`
  - `BufferBlock`-decorated `OpTypeStruct` in the `Uniform` storage `Storage` `Class`
  - the outermost dimension of an arrayed variable in the `StorageBuffer`, `Uniform`, or `UniformConstant` storage `Storage` `Class`
  - variables in the `NodePayloadAMDX` storage `Storage` `Class` when the `CoalescingAMDX` `Execution` `Mode` is specified
- [](#VUID-StandaloneSpirv-Function-04681)VUID-StandaloneSpirv-Function-04681  
  A type *T* that is an array sized with a specialization constant **must** neither be, nor be contained in, the type *T2* of a variable *V*, unless either: a) *T* is equal to *T2*, b) *V* is declared in the `Function`, or `Private` `Storage` `Class`, c) *V* is a non-Block variable in the `Workgroup` `Storage` `Class`, or d) *V* is an interface variable with an additional level of arrayness, [as described in interface matching](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#interfaces-iointerfaces-matching), and *T* is the member type of the array type *T2*
- [](#VUID-StandaloneSpirv-OpControlBarrier-04682)VUID-StandaloneSpirv-OpControlBarrier-04682  
  If `OpControlBarrier` is used in ray generation, intersection, any-hit, closest hit, miss, fragment, vertex, tessellation evaluation, or geometry shaders, the execution Scope **must** be `Subgroup`
- [](#VUID-StandaloneSpirv-None-10685)VUID-StandaloneSpirv-None-10685  
  For each compute shader entry point, either a `TileShadingRateQCOM`, `LocalSize`, or `LocalSizeId` `Execution` `Mode`, or an object decorated with the `WorkgroupSize` decoration **must** be specified
- [](#VUID-StandaloneSpirv-DerivativeGroupQuadsNV-04684)VUID-StandaloneSpirv-DerivativeGroupQuadsNV-04684  
  For compute shaders using the `DerivativeGroupQuadsNV` execution mode, the first two dimensions of the local workgroup size **must** be a multiple of two
- [](#VUID-StandaloneSpirv-DerivativeGroupLinearNV-04778)VUID-StandaloneSpirv-DerivativeGroupLinearNV-04778  
  For compute shaders using the `DerivativeGroupLinearNV` execution mode, the product of the dimensions of the local workgroup size **must** be a multiple of four
- [](#VUID-StandaloneSpirv-DerivativeGroupQuadsKHR-10151)VUID-StandaloneSpirv-DerivativeGroupQuadsKHR-10151  
  For compute, mesh, or task shaders using the `DerivativeGroupQuadsKHR` execution mode, the first two dimensions of the local workgroup size **must** be a multiple of two
- [](#VUID-StandaloneSpirv-DerivativeGroupLinearKHR-10152)VUID-StandaloneSpirv-DerivativeGroupLinearKHR-10152  
  For compute, mesh, or task shaders using the `DerivativeGroupLinearKHR` execution mode, the product of the dimensions of the local workgroup size **must** be a multiple of four
- [](#VUID-StandaloneSpirv-OpGroupNonUniformBallotBitCount-04685)VUID-StandaloneSpirv-OpGroupNonUniformBallotBitCount-04685  
  If `OpGroupNonUniformBallotBitCount` is used, the group operation **must** be limited to `Reduce`, `InclusiveScan`, or `ExclusiveScan`
- [](#VUID-StandaloneSpirv-None-04686)VUID-StandaloneSpirv-None-04686  
  The *Pointer* operand of all atomic instructions **must** have a `Storage` `Class` limited to `Uniform`, `Workgroup`, `Image`, `StorageBuffer`, `PhysicalStorageBuffer`, or `TaskPayloadWorkgroupEXT`
- [](#VUID-StandaloneSpirv-Offset-04687)VUID-StandaloneSpirv-Offset-04687  
  Output variables or block members decorated with `Offset` that have a 64-bit type, or a composite type containing a 64-bit type, **must** specify an `Offset` value aligned to a 8 byte boundary
- [](#VUID-StandaloneSpirv-Offset-04689)VUID-StandaloneSpirv-Offset-04689  
  The size of any output block containing any member decorated with `Offset` that is a 64-bit type **must** be a multiple of 8
- [](#VUID-StandaloneSpirv-Offset-04690)VUID-StandaloneSpirv-Offset-04690  
  The first member of an output block specifying a `Offset` decoration **must** specify a `Offset` value that is aligned to an 8 byte boundary if that block contains any member decorated with `Offset` and is a 64-bit type
- [](#VUID-StandaloneSpirv-Offset-04691)VUID-StandaloneSpirv-Offset-04691  
  Output variables or block members decorated with `Offset` that have a 32-bit type, or a composite type contains a 32-bit type, **must** specify an `Offset` value aligned to a 4 byte boundary
- [](#VUID-StandaloneSpirv-Offset-04692)VUID-StandaloneSpirv-Offset-04692  
  Output variables, blocks or block members decorated with `Offset` **must** only contain base types that have components that are either 32-bit or 64-bit in size
- [](#VUID-StandaloneSpirv-Offset-04716)VUID-StandaloneSpirv-Offset-04716  
  Only variables or block members in the output interface decorated with `Offset` **can** be captured for transform feedback, and those variables or block members **must** also be decorated with `XfbBuffer` and `XfbStride`, or inherit `XfbBuffer` and `XfbStride` decorations from a block containing them
- [](#VUID-StandaloneSpirv-XfbBuffer-04693)VUID-StandaloneSpirv-XfbBuffer-04693  
  All variables or block members in the output interface of the entry point being compiled decorated with a specific `XfbBuffer` value **must** all be decorated with identical `XfbStride` values
- [](#VUID-StandaloneSpirv-Stream-04694)VUID-StandaloneSpirv-Stream-04694  
  If any variables or block members in the output interface of the entry point being compiled are decorated with `Stream`, then all variables belonging to the same `XfbBuffer` **must** specify the same `Stream` value
- [](#VUID-StandaloneSpirv-XfbBuffer-04696)VUID-StandaloneSpirv-XfbBuffer-04696  
  For any two variables or block members in the output interface of the entry point being compiled with the same `XfbBuffer` value, the ranges determined by the `Offset` decoration and the size of the type **must** not overlap
- [](#VUID-StandaloneSpirv-XfbBuffer-04697)VUID-StandaloneSpirv-XfbBuffer-04697  
  All block members in the output interface of the entry point being compiled that are in the same block and have a declared or inherited `XfbBuffer` decoration **must** specify the same `XfbBuffer` value
- [](#VUID-StandaloneSpirv-RayPayloadKHR-04698)VUID-StandaloneSpirv-RayPayloadKHR-04698  
  `RayPayloadKHR` `Storage` `Class` **must** only be used in ray generation, closest hit or miss shaders
- [](#VUID-StandaloneSpirv-IncomingRayPayloadKHR-04699)VUID-StandaloneSpirv-IncomingRayPayloadKHR-04699  
  `IncomingRayPayloadKHR` `Storage` `Class` **must** only be used in closest hit, any-hit, or miss shaders
- [](#VUID-StandaloneSpirv-IncomingRayPayloadKHR-04700)VUID-StandaloneSpirv-IncomingRayPayloadKHR-04700  
  There **must** be at most one variable with the `IncomingRayPayloadKHR` `Storage` `Class` in the input interface of an entry point
- [](#VUID-StandaloneSpirv-HitAttributeKHR-04701)VUID-StandaloneSpirv-HitAttributeKHR-04701  
  `HitAttributeKHR` `Storage` `Class` **must** only be used in intersection, any-hit, or closest hit shaders
- [](#VUID-StandaloneSpirv-HitAttributeKHR-04702)VUID-StandaloneSpirv-HitAttributeKHR-04702  
  There **must** be at most one variable with the `HitAttributeKHR` `Storage` `Class` in the input interface of an entry point
- [](#VUID-StandaloneSpirv-HitAttributeKHR-04703)VUID-StandaloneSpirv-HitAttributeKHR-04703  
  A variable with `HitAttributeKHR` `Storage` `Class` **must** only be written to in an intersection shader
- [](#VUID-StandaloneSpirv-CallableDataKHR-04704)VUID-StandaloneSpirv-CallableDataKHR-04704  
  `CallableDataKHR` `Storage` `Class` **must** only be used in ray generation, closest hit, miss, and callable shaders
- [](#VUID-StandaloneSpirv-IncomingCallableDataKHR-04705)VUID-StandaloneSpirv-IncomingCallableDataKHR-04705  
  `IncomingCallableDataKHR` `Storage` `Class` **must** only be used in callable shaders
- [](#VUID-StandaloneSpirv-IncomingCallableDataKHR-04706)VUID-StandaloneSpirv-IncomingCallableDataKHR-04706  
  There **must** be at most one variable with the `IncomingCallableDataKHR` `Storage` `Class` in the input interface of an entry point
- [](#VUID-StandaloneSpirv-ShaderRecordBufferKHR-07119)VUID-StandaloneSpirv-ShaderRecordBufferKHR-07119  
  `ShaderRecordBufferKHR` `Storage` `Class` **must** only be used in ray generation, intersection, any-hit, closest hit, callable, or miss shaders
- [](#VUID-StandaloneSpirv-Base-07650)VUID-StandaloneSpirv-Base-07650  
  The `Base` operand of `OpPtrAccessChain` **must** have a storage class of `Workgroup`, `StorageBuffer`, or `PhysicalStorageBuffer`
- [](#VUID-StandaloneSpirv-Base-07651)VUID-StandaloneSpirv-Base-07651  
  If the `Base` operand of `OpPtrAccessChain` has a `Workgroup` `Storage` `Class`, then the `VariablePointers` capability **must** be declared
- [](#VUID-StandaloneSpirv-Base-07652)VUID-StandaloneSpirv-Base-07652  
  If the `Base` operand of `OpPtrAccessChain` has a `StorageBuffer` `Storage` `Class`, then the `VariablePointers` or `VariablePointersStorageBuffer` capability **must** be declared
- [](#VUID-StandaloneSpirv-PhysicalStorageBuffer64-04708)VUID-StandaloneSpirv-PhysicalStorageBuffer64-04708  
  If the `PhysicalStorageBuffer64` addressing model is enabled, all instructions that support memory access operands and that use a physical pointer **must** include the `Aligned` operand
- [](#VUID-StandaloneSpirv-PhysicalStorageBuffer64-04709)VUID-StandaloneSpirv-PhysicalStorageBuffer64-04709  
  If the `PhysicalStorageBuffer64` addressing model is enabled, any access chain instruction that accesses into a `RowMajor` matrix **must** only be used as the `Pointer` operand to `OpLoad` or `OpStore`
- [](#VUID-StandaloneSpirv-PhysicalStorageBuffer64-04710)VUID-StandaloneSpirv-PhysicalStorageBuffer64-04710  
  If the `PhysicalStorageBuffer64` addressing model is enabled, `OpConvertUToPtr` and `OpConvertPtrToU` **must** use an integer type whose `Width` is 64
- [](#VUID-StandaloneSpirv-OpTypeForwardPointer-04711)VUID-StandaloneSpirv-OpTypeForwardPointer-04711  
  `OpTypeForwardPointer` **must** have a `Storage` `Class` of `PhysicalStorageBuffer`
- [](#VUID-StandaloneSpirv-None-04745)VUID-StandaloneSpirv-None-04745  
  All block members in a variable with a `Storage` `Class` of `PushConstant` declared as an array **must** only be accessed by dynamically uniform indices
- [](#VUID-StandaloneSpirv-OpVariable-06673)VUID-StandaloneSpirv-OpVariable-06673  
  There **must** not be more than one `OpVariable` in the `PushConstant` `Storage` `Class` listed in the `Interface` for each `OpEntryPoint`
- [](#VUID-StandaloneSpirv-OpEntryPoint-06674)VUID-StandaloneSpirv-OpEntryPoint-06674  
  Each `OpEntryPoint` **must** not statically use more than one `OpVariable` in the `PushConstant` `Storage` `Class`
- [](#VUID-StandaloneSpirv-OpEntryPoint-08721)VUID-StandaloneSpirv-OpEntryPoint-08721  
  Each `OpEntryPoint` **must** not have more than one `Input` variable assigned the same `Component` word inside a `Location` slot, either explicitly or implicitly
- [](#VUID-StandaloneSpirv-OpEntryPoint-08722)VUID-StandaloneSpirv-OpEntryPoint-08722  
  Each `OpEntryPoint` **must** not have more than one `Output` variable assigned the same `Component` word inside a `Location` slot, either explicitly or implicitly
- [](#VUID-StandaloneSpirv-Result-04780)VUID-StandaloneSpirv-Result-04780  
  The `Result` `Type` operand of any `OpImageRead` or `OpImageSparseRead` instruction **must** be a vector of four components
- [](#VUID-StandaloneSpirv-PushConstant-06675)VUID-StandaloneSpirv-PushConstant-06675  
  Any variable in the `PushConstant` or `StorageBuffer` storage class **must** be decorated as `Block`
- [](#VUID-StandaloneSpirv-Uniform-06676)VUID-StandaloneSpirv-Uniform-06676  
  Any variable in the `Uniform` `Storage` `Class` **must** be decorated as `Block` or `BufferBlock`
- [](#VUID-StandaloneSpirv-UniformConstant-06677)VUID-StandaloneSpirv-UniformConstant-06677  
  Any variable in the `UniformConstant`, `StorageBuffer`, or `Uniform` `Storage` `Class` **must** be decorated with `DescriptorSet` and `Binding`
- [](#VUID-StandaloneSpirv-InputAttachmentIndex-06678)VUID-StandaloneSpirv-InputAttachmentIndex-06678  
  Variables decorated with `InputAttachmentIndex` **must** be in the `UniformConstant` `Storage` `Class`
- [](#VUID-StandaloneSpirv-DescriptorSet-06491)VUID-StandaloneSpirv-DescriptorSet-06491  
  If a variable is decorated by `DescriptorSet` or `Binding`, the `Storage` `Class` **must** correspond to an entry in [Shader Resource and Storage Class Correspondence](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#interfaces-resources-storage-class-correspondence)
- [](#VUID-StandaloneSpirv-Input-06778)VUID-StandaloneSpirv-Input-06778  
  Variables with a `Storage` `Class` of `Input` in a fragment shader stage that are decorated with `PerVertexKHR` **must** be declared as arrays
- [](#VUID-StandaloneSpirv-MeshEXT-07102)VUID-StandaloneSpirv-MeshEXT-07102  
  The module **must** not contain both an entry point that uses the `TaskEXT` or `MeshEXT` `Execution` `Model` and an entry point that uses the `TaskNV` or `MeshNV` `Execution` `Model`
- [](#VUID-StandaloneSpirv-MeshEXT-07106)VUID-StandaloneSpirv-MeshEXT-07106  
  In mesh shaders using the `MeshEXT` `Execution` `Model` `OpSetMeshOutputsEXT` **must** be called before any outputs are written
- [](#VUID-StandaloneSpirv-MeshEXT-07107)VUID-StandaloneSpirv-MeshEXT-07107  
  In mesh shaders using the `MeshEXT` `Execution` `Model` all variables declared as output **must** not be read from
- [](#VUID-StandaloneSpirv-MeshEXT-07108)VUID-StandaloneSpirv-MeshEXT-07108  
  In mesh shaders using the `MeshEXT` `Execution` `Model` for `OpSetMeshOutputsEXT` instructions, the “Vertex Count” and “Primitive Count” operands **must** not depend on `ViewIndex`
- [](#VUID-StandaloneSpirv-MeshEXT-07109)VUID-StandaloneSpirv-MeshEXT-07109  
  In mesh shaders using the `MeshEXT` `Execution` `Model` variables decorated with `PrimitivePointIndicesEXT`, `PrimitiveLineIndicesEXT`, or `PrimitiveTriangleIndicesEXT` declared as an array **must** not be accessed by indices that depend on `ViewIndex`
- [](#VUID-StandaloneSpirv-MeshEXT-07110)VUID-StandaloneSpirv-MeshEXT-07110  
  In mesh shaders using the `MeshEXT` `Execution` `Model` any values stored in variables decorated with `PrimitivePointIndicesEXT`, `PrimitiveLineIndicesEXT`, or `PrimitiveTriangleIndicesEXT` **must** not depend on `ViewIndex`
- [](#VUID-StandaloneSpirv-MeshEXT-07111)VUID-StandaloneSpirv-MeshEXT-07111  
  In mesh shaders using the `MeshEXT` `Execution` `Model` variables in workgroup or private `Storage` `Class` declared as or containing a composite type **must** not be accessed by indices that depend on `ViewIndex`
- [](#VUID-StandaloneSpirv-MeshEXT-07330)VUID-StandaloneSpirv-MeshEXT-07330  
  In mesh shaders using the `MeshEXT` `Execution` `Model` the `OutputVertices` `Execution` `Mode` **must** be greater than 0
- [](#VUID-StandaloneSpirv-MeshEXT-07331)VUID-StandaloneSpirv-MeshEXT-07331  
  In mesh shaders using the `MeshEXT` `Execution` `Model` the `OutputPrimitivesEXT` `Execution` `Mode` **must** be greater than 0
- [](#VUID-StandaloneSpirv-Input-07290)VUID-StandaloneSpirv-Input-07290  
  Variables with a `Storage` `Class` of `Input` or `Output` and a type of `OpTypeBool` **must** be decorated with the `BuiltIn` decoration
- [](#VUID-StandaloneSpirv-TileImageEXT-08723)VUID-StandaloneSpirv-TileImageEXT-08723  
  The tile image variable declarations **must** obey the constraints on the `TileImageEXT` `Storage` `Class` and the `Location` decoration described in [Fragment Tile Image Interface](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#interfaces-fragmenttileimage)
- [](#VUID-StandaloneSpirv-None-08724)VUID-StandaloneSpirv-None-08724  
  The `TileImageEXT` `Storage` `Class` **must** only be used for declaring tile image variables
- [](#VUID-StandaloneSpirv-Pointer-08973)VUID-StandaloneSpirv-Pointer-08973  
  The `Storage` `Class` of the `Pointer` operand to `OpCooperativeMatrixLoadKHR` or `OpCooperativeMatrixStoreKHR` **must** be limited to `Workgroup`, `StorageBuffer`, or `PhysicalStorageBuffer`
- [](#VUID-StandaloneSpirv-OpTypeFloat-10370)VUID-StandaloneSpirv-OpTypeFloat-10370  
  Variables with a type of `OpTypeFloat` and an encoding of `BFloat16KHR` **must** not be declared with a `Storage` `Class` of `Input` or `Output`
- [](#VUID-StandaloneSpirv-UniformBufferArrayDynamicIndexing-10127)VUID-StandaloneSpirv-UniformBufferArrayDynamicIndexing-10127  
  If the `UniformBufferArrayDynamicIndexing` capability is not declared, and an instruction accesses memory through a uniform buffer, the uniform buffer through which that memory is accessed **must** be determined by [constant integral expressions](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#glossary-constant-integral-expression)
- [](#VUID-StandaloneSpirv-SampledImageArrayDynamicIndexing-10128)VUID-StandaloneSpirv-SampledImageArrayDynamicIndexing-10128  
  If the `SampledImageArrayDynamicIndexing` capability is not declared, and an instruction accesses memory through a sampled image or sampler, the sampled image or sampler through which that memory is accessed **must** be determined by [constant integral expressions](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#glossary-constant-integral-expression)
- [](#VUID-StandaloneSpirv-StorageBufferArrayDynamicIndexing-10129)VUID-StandaloneSpirv-StorageBufferArrayDynamicIndexing-10129  
  If the `StorageBufferArrayDynamicIndexing` capability is not declared, and an instruction accesses memory through a storage buffer, the storage buffer through which that memory is accessed **must** be determined by [constant integral expressions](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#glossary-constant-integral-expression)
- [](#VUID-StandaloneSpirv-StorageImageArrayDynamicIndexing-10130)VUID-StandaloneSpirv-StorageImageArrayDynamicIndexing-10130  
  If the `StorageImageArrayDynamicIndexing` capability is not declared, and an instruction accesses memory through a storage image, the storage image through which that memory is accessed **must** be determined by [constant integral expressions](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#glossary-constant-integral-expression)
- [](#VUID-StandaloneSpirv-InputAttachmentArrayDynamicIndexing-10131)VUID-StandaloneSpirv-InputAttachmentArrayDynamicIndexing-10131  
  If the `InputAttachmentArrayDynamicIndexing` capability is not declared, and an instruction accesses memory through an input attachment, the input attachmnet through which that memory is accessed **must** be determined by [constant integral expressions](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#glossary-constant-integral-expression)
- [](#VUID-StandaloneSpirv-UniformTexelBufferArrayDynamicIndexing-10132)VUID-StandaloneSpirv-UniformTexelBufferArrayDynamicIndexing-10132  
  If the `UniformTexelBufferArrayDynamicIndexing` capability is not declared, and an instruction accesses memory through a uniform texel buffer, the uniform texel buffer through which that memory is accessed **must** be determined by [constant integral expressions](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#glossary-constant-integral-expression)
- [](#VUID-StandaloneSpirv-StorageTexelBufferArrayDynamicIndexing-10133)VUID-StandaloneSpirv-StorageTexelBufferArrayDynamicIndexing-10133  
  If the `StorageTexelBufferArrayDynamicIndexing` capability is not declared, and an instruction accesses memory through a storage texel buffer, the storage texel buffer through which that memory is accessed **must** be determined by [constant integral expressions](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#glossary-constant-integral-expression)
- [](#VUID-StandaloneSpirv-TileShadingQCOM-10686)VUID-StandaloneSpirv-TileShadingQCOM-10686  
  `TileShadingQCOM` capability **must** not be enabled in any stage other than compute or fragment
- [](#VUID-StandaloneSpirv-Execution-10687)VUID-StandaloneSpirv-Execution-10687  
  `Execution` `Mode` `NonCoherentTileAttachmentReadQCOM` **must** not be used in any stage other than fragment
- [](#VUID-StandaloneSpirv-Execution-10688)VUID-StandaloneSpirv-Execution-10688  
  `Execution` `Mode` `TileShadingRateQCOM` **must** not be used in any stage other than compute
- [](#VUID-StandaloneSpirv-TileAttachmentQCOM-10689)VUID-StandaloneSpirv-TileAttachmentQCOM-10689  
  If the `TileAttachmentQCOM` `Storage` `Class` is used, the `TileShadingQCOM` capability **must** be enabled
- [](#VUID-StandaloneSpirv-NonCoherentTileAttachmentReadQCOM-10690)VUID-StandaloneSpirv-NonCoherentTileAttachmentReadQCOM-10690  
  If the `NonCoherentTileAttachmentReadQCOM` `Execution` `Mode` is used, the `TileShadingQCOM` capability **must** be enabled
- [](#VUID-StandaloneSpirv-TileShadingRateQCOM-10691)VUID-StandaloneSpirv-TileShadingRateQCOM-10691  
  If the `TileShadingRateQCOM` `Execution` `Mode` is used, the `TileShadingQCOM` capability **must** be enabled
- [](#VUID-StandaloneSpirv-TileShadingRateQCOM-10692)VUID-StandaloneSpirv-TileShadingRateQCOM-10692  
  If the `TileShadingRateQCOM` `Execution` `Mode` is used, `LocalSize` and `LocalSizeId` `Execution` `Mode` **must** not be specified
- [](#VUID-StandaloneSpirv-OpTypeImage-10693)VUID-StandaloneSpirv-OpTypeImage-10693  
  `OpTypeImage` variables in the `TileAttachmentQCOM` `Storage` `Class` **must** have `Dim` equal to `2D`
- [](#VUID-StandaloneSpirv-OpTypeImage-10694)VUID-StandaloneSpirv-OpTypeImage-10694  
  `OpTypeImage` variables in the `TileAttachmentQCOM` `Storage` `Class` **must** `Sampled` equal to `1` or `2`
- [](#VUID-StandaloneSpirv-TileAttachmentQCOM-10695)VUID-StandaloneSpirv-TileAttachmentQCOM-10695  
  Any variable in the `TileAttachmentQCOM` `Storage` `Class` **must** be decorated with `DescriptorSet` and `Binding`
- [](#VUID-StandaloneSpirv-TileAttachmentQCOM-10696)VUID-StandaloneSpirv-TileAttachmentQCOM-10696  
  Any variable in the `TileAttachmentQCOM` `Storage` `Class` **must** not be decorated with `Component` decoration
- [](#VUID-StandaloneSpirv-TileAttachmentQCOM-10697)VUID-StandaloneSpirv-TileAttachmentQCOM-10697  
  An OpTypeImage variables in the `TileAttachmentQCOM` `Storage` `Class` **must** not be consumed by an `OpImageQuery*` instruction
- [](#VUID-StandaloneSpirv-OpTypeFloat-10823)VUID-StandaloneSpirv-OpTypeFloat-10823  
  Variables with a type of `OpTypeFloat` and an encoding of `Float8E4M3EXT` or `Float8E5M2EXT` **must** not be declared with a `Storage` `Class` of `Input` or `Output`
- [](#VUID-StandaloneSpirv-OpGraphInputARM-09931)VUID-StandaloneSpirv-OpGraphInputARM-09931  
  The `InputIndex` and `ElementIndex` operands to `OpGraphInputARM` **must** be the &lt;id&gt; of a constant instruction
- [](#VUID-StandaloneSpirv-OpGraphSetOutputARM-09932)VUID-StandaloneSpirv-OpGraphSetOutputARM-09932  
  The `OutputIndex` and `ElementIndex` operands to `OpGraphSetOutputARM` **must** be the &lt;id&gt; of a constant instruction

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#StandaloneSpirv)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0